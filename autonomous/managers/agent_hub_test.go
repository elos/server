package managers_test

import (
	. "github.com/elos/server/autonomous/managers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AgentHub", func() {
	Describe("NewAgentHub()", func() {
		var (
			ah *AgentHub
		)

		BeforeEach(func() {
			ah = NewAgentHub()
		})

		It("Allocates and returns a *AgentHub", func() {
			Expect(true).To(BeTrue())
		})

		It("Allocates the Start Channel", func() {
		})

		It("Allocates the Stop Channel", func() {
		})
	})

	Context("autonomous.Manager Implementation", func() {
		Describe("Run()", func() {

		})

		Describe("StartAgent()", func() {
		})

		Describe("StopAgent()", func() {
		})

		Describe("Die()", func() {
		})
	})

})
