package conn_test

import (
	. "github.com/elos/server/conn"
	"github.com/elos/server/data/models/user"

	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("GorillaConnection", func() {
	u := user.New()
	c := NewNullConnection(u)

	gc := NewGorillaConnection(c, u)

	Describe("NewGorillaConnection", func() {
		It("Allocates and returns a new Connection", func() {
			Expect(gc).NotTo(BeNil())
			Expect(gc).To(BeAssignableToTypeOf(&GorillaConnection{}))
		})
	})

	Describe("WriteJSON", func() {
		It("Passes WriteJSON to underlying AnonConnection", func() {
			err := gc.WriteJSON(u)
			Expect(err).NotTo(HaveOccurred())
			Expect(c.Writes).To(HaveKeyWithValue(u, true))

			c.Reset()
		})

		It("Will return an error", func() {
			e := errors.New("error")
			c.Error = e
			err := gc.WriteJSON(u)

			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(e))
			Expect(c.Writes).To(BeEmpty())

			c.Reset()
		})
	})

	Describe("ReadJSON", func() {
		It("Passes ReadJSON to underlying AnonConnection", func() {
			err := gc.ReadJSON(u)

			Expect(err).NotTo(HaveOccurred())
			Expect(c.Reads).To(HaveKeyWithValue(u, true))

			c.Reset()
		})

		It("Returns an error if needed", func() {
			e := errors.New("error")
			c.Error = e
			err := gc.WriteJSON(u)

			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(e))
			Expect(c.Reads).To(BeEmpty())

			c.Reset()
		})
	})

	Describe("Close()", func() {
		It("Passes the close to the underlying AnonConnection", func() {
			gc.Close()
			Expect(c.Closed).To(BeTrue())
			c.Reset()
		})
	})

	Describe("Agent()", func() {
		It("Returns the agent that was passed in on creation", func() {
			a := gc.Agent()
			Expect(a).To(Equal(u))
		})
	})

})
