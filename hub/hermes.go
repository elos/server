package hub

import "gopkg.in/mgo.v2/bson"

var PrimaryHermes *Hermes

type Package interface {
	// Concerned are the recipients of the package
	Concerned() []*bson.ObjectId
}

type ConcernedUpdate struct {
	ConcernedId *bson.ObjectId
	Package     Package
}

type Hermes struct {
	Send chan Package
}

func SetupHermes() {
	PrimaryHermes = &Hermes{
		Send: make(chan Package),
	}

	PrimaryHermes.Run()
}

func (h *Hermes) Run() {
	for {
		select {
		case p := <-h.Send:
			recipientIds := p.Concerned()

			for _, recipientId := range recipientIds {
				update := &ConcernedUpdate{
					ConcernedId: recipientId,
					Package:     p,
				}

				PrimaryHub.Update <- update
			}
		}
	}
}
