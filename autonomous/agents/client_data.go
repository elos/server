package agents

import (
	"github.com/elos/server/autonomous"
	"github.com/elos/server/conn"
	"github.com/elos/server/data"
	"github.com/elos/server/data/transfer"
	"log"
)

type ClientDataAgent struct {
	running bool
	stop    chan bool
	read    chan *data.Envelope

	DB         data.DB
	DataAgent  data.Agent
	Manager    autonomous.Manager
	Connection conn.Connection
}

func NewClientDataAgent(c conn.Connection, db data.DB) autonomous.Agent {
	a := &ClientDataAgent{
		Connection: c,
		DB:         db,
		stop:       make(chan bool),
		read:       make(chan *data.Envelope),
	}

	a.SetDataAgent(c.Agent())

	return a
}

func (a *ClientDataAgent) SetDataAgent(da data.Agent) {
	if da != nil {
		a.DataAgent = da
	}
}

func (a *ClientDataAgent) GetDataAgent() data.Agent {
	return a.DataAgent
}

func (a *ClientDataAgent) SetManager(m autonomous.Manager) {
	if m != nil {
		a.Manager = m
	}
}

func (a *ClientDataAgent) GetManager() autonomous.Manager {
	return a.Manager
}

func (a *ClientDataAgent) Start() {
	go ReadConnection(a.Connection, &a.read, &a.stop)

	modelsChannel := *a.DB.RegisterForUpdates(a.DataAgent)

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

func (a *ClientDataAgent) Stop() {
	a.stop <- true
}

func (a *ClientDataAgent) Kill() {
	a.stop <- true
}

func (a *ClientDataAgent) Alive() bool {
	return a.running
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
