package sockets_test

import (
	. "github.com/elos/server/sockets"

	"github.com/elos/server/db"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hub", func() {
	Context("Hub Type", func() {
		It("Should exist", func() {
			h := Hub{}
			Expect(h).ToNot(BeNil())
		})

		It("NewHub()", func() {
			channel := make(chan db.Model)
			db := &db.MongoDB{ModelUpdates: &channel}
			h := NewHub(db)
			Expect(h.Channels).ToNot(BeNil())
			Expect(h.Register).ToNot(BeNil())
			Expect(h.Unregister).ToNot(BeNil())
		})
	})

	Context("Setup", func() {
		BeforeEach(func() {
			Expect(PrimaryHub).To(BeNil())
			channel := make(chan db.Model)
			Setup(&db.MongoDB{ModelUpdates: &channel})
		})

		It("Should set the Primary Hub", func() {
			Expect(PrimaryHub).ToNot(BeNil())
		})
	})

	Context("Connection Registering", func() {
		var (
			h *Hub
		)

		channel := make(chan db.Model)
		db := &db.MongoDB{ModelUpdates: &channel}

		BeforeEach(func() {
			h = NewHub(db)
			go h.Run()
		})

		Describe("RegisterConnection", func() {
			It("Registers a connection", func() {
				// need to test connections first :/
			})
		})

		Describe("UnregisterConnection", func() {
			It("Unregisters a connection", func() {
			})
		})
	})
})
