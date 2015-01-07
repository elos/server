package routes_test

import (
	. "github.com/elos/server/routes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("WebSocket", func() {

	It("Defines a Default websocket upgrader", func() {
		Expect(DefaultWebSocketUpgrader).NotTo(BeNil())
	})

	It("Defines a WebSocketProtocolHeader", func() {
		Expect(WebSocketProtocolHeader).NotTo(BeNil())
	})

	Describe("ExtractProtocolHeader", func() {
		p := "askldfjasdjfkjalsdfljkasdjkfasdflkaf"

		header := http.Header{}
		header.Add(WebSocketProtocolHeader, p)

		r := &http.Request{
			Header: header,
		}

		It("Extracts the websocket protocol header into a new header objects", func() {
			h := ExtractProtocolHeader(r)
			Expect(h).NotTo(BeNil())
			Expect(h).To(BeAssignableToTypeOf(http.Header{}))
			Expect(h.Get(WebSocketProtocolHeader)).To(Equal(p))
		})

	})

})
