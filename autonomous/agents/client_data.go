package agents

import (
	"github.com/elos/server/autonomous"
	"github.com/elos/server/conn"
	"github.com/elos/server/data"
	"github.com/elos/server/data/transfer"
	"log"
)

type ClientDataAgent struct {
	*BaseAgent
	DB data.DB

	read       chan *data.Envelope
	Connection conn.Connection
}

func NewClientDataAgent(c conn.Connection, db data.DB) autonomous.Agent {
	a := &ClientDataAgent{
		BaseAgent:  NewBaseAgent(),
		Connection: c,
		DB:         db,
		read:       make(chan *data.Envelope),
	}

	a.SetDataAgent(c.Agent())

	return a
}

func (a *ClientDataAgent) Start() {
	go ReadConnection(a.Connection, &a.read, &a.stop)

	modelsChannel := *a.DB.RegisterForUpdates(a.GetDataAgent())

	for {
		select {
		case e := <-a.read:
			log.Print("WE HAVE A READ")
			go transfer.Route(e, a.DB, a.Connection)
			continue
		case p := <-modelsChannel:
			log.Print("WE HAVE AN UPDATE")
			a.WriteJSON(p)
		case _ = <-a.stop:
			//shutdown
			continue
		}
	}
}

func ReadConnection(c conn.Connection, rc *chan *data.Envelope, endChannel *chan bool) {
	// TODO add read limit and deadline
	for {
		var e data.Envelope

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
