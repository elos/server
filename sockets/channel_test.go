package sockets_test

import (
	. "github.com/elos/server/sockets"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Channel", func() {
	Context("Type Declaration", func() {
		Describe("Channel Type", func() {
			It("Exists", func() {
				var channel *Channel
				channel = &Channel{
					Connections: make([]*Connection, 1),
					Send:        make(chan []byte),
				}
				Expect(channel).ToNot(BeNil())
				Expect(channel.Connections).ToNot(BeNil())
				Expect(channel.Connections).To(HaveLen(1))
				Expect(channel.Send).ToNot(BeNil())
			})
		})

		Describe("Channel Creation", func() {
			It("Creates a new channel with empty connections array", func() {
				c := NewChannel()
				Expect(c).ToNot(BeNil())
				Expect(c.Connections).To(HaveLen(0))
			})
		})
	})

	Context("Generic Operations", func() {
		Describe("AddConnection", func() {
		})
		Describe("RemoveConnection", func() {
		})
		Describe("IndexConnection", func() {
		})
		Describe("DeleteConnection", func() {
		})
	})
})
