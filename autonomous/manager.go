package autonomous

type Manager interface {
	Run()

	StartAgent(Agent)
	StopAgent(Agent)

	Die()
}
