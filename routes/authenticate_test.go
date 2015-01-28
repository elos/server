package routes_test

import (
	"github.com/elos/server/conn"
	. "github.com/elos/server/routes"

	"net/http"
	"net/http/httptest"

	"github.com/elos/server/models/user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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
