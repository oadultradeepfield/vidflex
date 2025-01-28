package simulator

import (
	"sync"
)

type Simulator struct {
	Sessions           map[string]*Session
	Mu                 sync.RWMutex
	DipProbability     float64
	DipDurationInTicks int
}

func NewSimulator() *Simulator {
	return &Simulator{
		Sessions:           make(map[string]*Session),
		DipProbability:     0.1,
		DipDurationInTicks: 3,
	}
}
