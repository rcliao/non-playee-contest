package contest

import "fmt"

// Environment represents a overall problem for Agents to play with.
// It needs to receives input from agents and process inputs to change its
// internal state
type Environment interface {
	Start(config string) error
	Turn(input []string) (string, error)
	GetState() string
	IsEnded() bool
}

// Agent represents individual submission from the user
type Agent interface {
	Start(config string) error
	Turn(input string) (string, error)
}

// Manage is the top level life cycle between environment and its agents
func Manage(environment Environment, agents []Agent) {
	for environment.IsEnded() {
		state := environment.GetState()
		moves := []string{}
		for _, agent := range agents {
			agentMove, err := agent.Turn(state)
			if err != nil {
				// mark agent as dead and move onto next agent
				fmt.Println(err)
			}
			moves = append(moves, agentMove)
		}
		environment.Turn(moves)
	}
}
