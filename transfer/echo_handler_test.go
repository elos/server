package transfer_test

import (
	. "github.com/elos/server/data/transfer"

	"github.com/elos/server/conn"
	"github.com/elos/server/data"
	"github.com/elos/server/data/models/user"
	"github.com/elos/server/data/test"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EchoHandler", func() {
	It("Responds with the exact envelope given it", func() {
		a := user.New()
		c := conn.NewNullConnection(a)
		db := test.NewDB()
		id := data.NewObjectID()

		info := map[data.Kind]map[string]interface{}{
			user.Kind: {
				"id": id.Hex(),
			},
		}

		e := data.NewEnvelope(POST, info)
		EchoHandler(e, db, c)

		Expect(c.LastWrite).To(Equal(e))
	})
})
