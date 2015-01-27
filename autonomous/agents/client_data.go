package agents

import (
	"github.com/elos/data"
	"github.com/elos/server/autonomous"
	"github.com/elos/server/conn"
	"github.com/elos/server/transfer"
	"log"
)

type ClientDataAgent struct {
	*autonomous.BaseAgent
	DB data.DB

	read       chan *transfer.Envelope
	Connection conn.Connection
}

func NewClientDataAgent(c conn.Connection, db data.DB) autonomous.Agent {
	a := &ClientDataAgent{
		BaseAgent:  autonomous.NewBaseAgent(),
		Connection: c,
		DB:         db,
		read:       make(chan *transfer.Envelope),
	}

	a.SetDataOwner(c.Agent())

	return a
}

func (a *ClientDataAgent) Start() {
	a.startup()
	stopChannel := a.BaseAgent.StopChannel()
	modelsChannel := *a.DB.RegisterForUpdates(a.GetDataOwner())

	for {
		select {
		case e := <-a.read:
			log.Print("WE HAVE A READ")
			go transfer.Route(e, a.DB, a.Connection)
			continue
		case p := <-modelsChannel:
			log.Print("WE HAVE AN UPDATE")
			a.WriteJSON(p)
		case _ = <-*stopChannel:
			a.shutdown()
			break
		}
	}
}

func (a *ClientDataAgent) startup() {
	a.BaseAgent.Startup()
	go ReadConnection(a.Connection, &a.read, a.BaseAgent.StopChannel())
}

func (a *ClientDataAgent) shutdown() {
	a.BaseAgent.Shutdown()
}

func ReadConnection(c conn.Connection, rc *chan *transfer.Envelope, endChannel *chan bool) {
	// TODO add read limit and deadline
	for {
		var e transfer.Envelope

		err := c.ReadJSON(&e)

		if err != nil {
			//Logf("An error occurred while reading a Connection, err: %s", err)

			/*
				If there was an error break inf. loop.
				Function then completes, and endChannel is called
			*/
			break
		}

		*rc <- &e
	}

	*endChannel <- true
}

func (a *ClientDataAgent) WriteJSON(v interface{}) {
	a.Connection.WriteJSON(v)
}
