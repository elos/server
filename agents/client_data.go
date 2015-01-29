package agents

import (
	"log"

	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/server/conn"
	"github.com/elos/server/transfer"
)

type ClientDataAgent struct {
	*autonomous.Core
	*autonomous.Identified
	data.Store

	read       chan *transfer.Envelope
	Connection conn.Connection
}

func NewClientDataAgent(c conn.Connection, s data.Store) autonomous.Agent {
	a := &ClientDataAgent{
		Core:       autonomous.NewCore(),
		Identified: autonomous.NewIdentified(),
		Connection: c,
		Store:      s,
		read:       make(chan *transfer.Envelope),
	}

	a.SetDataOwner(c.Agent())

	return a
}

func (a *ClientDataAgent) Run() {
	a.startup()
	stopChannel := a.Core.StopChannel()
	modelsChannel := *a.Store.RegisterForUpdates(a.DataOwner())

	for {
		select {
		case e := <-a.read:
			log.Print("WE HAVE A READ")
			go transfer.Route(e, a.Store, a.Connection)
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
	a.Core.Startup()
	go ReadConnection(a.Connection, &a.read, a.Core.StopChannel())
}

func (a *ClientDataAgent) shutdown() {
	a.Core.Shutdown()
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
