package routes_test

import (
	"github.com/elos/server/conn"
	. "github.com/elos/server/routes"

	"github.com/elos/server/data/models/user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Authenticate", func() {
	It("Exposes an AuthenticatedGet Handler", func() {
		Expect(AuthenticateGet).NotTo(BeNil())
	})

	Describe("WebSocketUpgradeHandler", func() {
		w := httptest.NewRecorder()
		r := &http.Request{}

		u := user.New()

		WebSocketUpgradeHandler(w, r, u, conn.DefaultWebSocketUpgrader, DefaultClientDataHub)
	})
})
