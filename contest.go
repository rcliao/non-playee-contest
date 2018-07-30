package contest

import "fmt"

type Environment interface {
	Start(config string) error
	Turn(input []string) (string, error)
	GetState() string
	IsEnded() bool
}

type Agent interface {
	Start(config string) error
	Turn(input string) (string, error)
}

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
