package hub

import (
	"log"

	"github.com/gorilla/websocket"
)

type HubChannel struct {
	Sockets []*websocket.Conn

	Send chan []byte
}

func (hc *HubChannel) AddSocket(s *websocket.Conn) {
	hc.Sockets = append(hc.Sockets, s)
}

func (hc *HubChannel) RemoveSocket(s *websocket.Conn) {
	DeleteSocket(hc.Sockets, s)
}

func IndexSocket(slice []*websocket.Conn, value *websocket.Conn) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func DeleteSocket(slice []*websocket.Conn, value *websocket.Conn) {
	i := IndexSocket(slice, value)
	a := slice
	a = a[:i+copy(a[i:], a[i+1:])]
}

func (hc *HubChannel) WriteJson(v interface{}) []error {
	log.Print("HUB channel write json", *hc)
	log.Print("hc.Sockets", hc.Sockets)

	var errs []error

	for _, socket := range hc.Sockets {
		err := socket.WriteJSON(v)

		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}
