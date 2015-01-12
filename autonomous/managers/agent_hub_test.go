package managers_test

import (
	"github.com/elos/server/autonomous/agents"
	. "github.com/elos/server/autonomous/managers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AgentHub", func() {
	var (
		ah *AgentHub
	)

	BeforeEach(func() {
		ah = NewAgentHub()
	})

	Describe("NewAgentHub()", func() {

		It("Allocates and returns a *AgentHub", func() {
			Expect(ah).NotTo(BeNil())
			Expect(ah).To(BeAssignableToTypeOf(&AgentHub{}))
		})

		It("Allocates the Start Channel", func() {
			Expect(ah.Start).NotTo(BeNil())
			Expect(ah.Start).To(BeEmpty())
		})

		It("Allocates the Stop Channel", func() {
			Expect(ah.Stop).NotTo(BeNil())
			Expect(ah.Stop).To(BeEmpty())
		})
	})

	Context("autonomous.Manager Implementation", func() {

		var (
			a *agents.BaseAgent
		)

		BeforeEach(func() {
			a = agents.NewBaseAgent()
		})

		Describe("Run()", func() {
			It("Returns nothing and starts listening on it's channels", func() {
				go ah.Run()
				Expect(ah.Start).To(Receive())
				Expect(ah.Stop).To(Receive())
			})
		})

		Describe("StartAgent()", func() {
			It("Starts the agent", func() {
				go ah.Run()
				ah.StartAgent(a)
				Expect(a.Alive()).To(BeTrue())
			})
		})

		Describe("StopAgent()", func() {
			It("Stops the agent", func() {
				go ah.Run()
				ah.StartAgent(a)
				ah.StopAgent(a)
				Expect(a.Alive()).To(BeTrue())
			})
		})

		Describe("Die()", func() {
			It("Stops all the agents and shuts down", func() {
				a2 := agents.NewBaseAgent()
				ah.StartAgent(a)
				ah.StartAgent(a2)
				ah.Die()

				Expect(a.Alive()).To(BeFalse())
				Expect(a2.Alive()).To(BeFalse())
			})
		})
	})

})
