package sockets_test

import (
	. "github.com/elos/server/sockets"

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
			h := NewHub()
			Expect(h.Channels).ToNot(BeNil())
			Expect(h.Register).ToNot(BeNil())
			Expect(h.Unregister).ToNot(BeNil())
		})
	})

	Context("Setup", func() {
		BeforeEach(func() {
			Expect(PrimaryHub).To(BeNil())
			Setup()
		})

		It("Should set the Primary Hub", func() {
			Expect(PrimaryHub).ToNot(BeNil())
		})
	})

})
