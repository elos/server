package managers_test

import (
	"github.com/elos/server/autonomous/agents"
	. "github.com/elos/server/autonomous/managers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NullHub", func() {

	var (
		n *NullHub
	)

	BeforeEach(func() {
		n = NewNullHub()
	})

	Describe("NewNullHub()", func() {
		It("Allocates and returns a new *NullHub", func() {
			Expect(n).NotTo(BeNil())
			Expect(n).To(BeAssignableToTypeOf(&NullHub{}))
			Expect(n.RegisteredAgents).NotTo(BeNil())
			Expect(n.RegisteredAgents).To(BeEmpty())
		})
	})

	Describe("StartAgent()", func() {
		It("Adds the agent to RegisteredAgent", func() {
			defer n.Reset()

			a := &agents.ClientDataAgent{}
			n.StartAgent(a)
			Expect(n.RegisteredAgents).To(HaveKeyWithValue(a, true))
		})
	})

	Describe("StopAgent()", func() {
		It("Set's the agent's RegisteredAgents key to be false", func() {
			defer n.Reset()

			a := &agents.ClientDataAgent{}
			n.StartAgent(a)
			n.StopAgent(a)
			Expect(n.RegisteredAgents).To(HaveKeyWithValue(a, false))
		})
	})

	Describe("Run()", func() {
		It("Sets alive to true", func() {
			defer n.Reset()

			n.Run()
			Expect(n.Alive).To(BeTrue())
		})
	})

	Describe("Die()", func() {
		It("Sets alive to false", func() {
			defer n.Reset()

			n.Die()
			Expect(n.Alive).To(BeFalse())
		})
	})

	Describe("Reset()", func() {
		It("Clears necessary fields", func() {
			n.Run()
			a := &agents.ClientDataAgent{}
			n.StartAgent(a)
			n.Reset()
			By("Setting RegisteredAgents to a newly initialized map")
			Expect(n.RegisteredAgents).To(BeEmpty())
			By("Setting Alive to false")
			Expect(n.Alive).To(BeFalse())
		})
	})
})
