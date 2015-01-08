package conn_test

import (
	. "github.com/elos/server/conn"

	"github.com/elos/server/data/models/user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NullConnection", func() {

	u := user.New()
	c := NewNullConnection(u)

	Describe("NewNullConnection", func() {
		It("Allocates and returns a new NullConnection", func() {
			Expect(c).NotTo(BeNil())
			Expect(c).To(BeAssignableToTypeOf(&NullConnection{}))
		})

		It("Allocates Writes and Reads maps", func() {
			Expect(c.Writes).NotTo(BeNil())
			Expect(c.Reads).NotTo(BeNil())
			Expect(c.Writes).To(BeEmpty())
			Expect(c.Reads).To(BeEmpty())
		})
	})

})
