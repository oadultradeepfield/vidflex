package simulator

import (
	"sync"
	"time"
)

type Simulator struct {
	Sessions       map[string]*Session
	DipProbability float64
	DipDuration    time.Duration
	Mu             sync.RWMutex
}

func NewSimulator() *Simulator {
	return &Simulator{
		Sessions:       make(map[string]*Session),
		DipProbability: 0.1,
		DipDuration:    3 * time.Second,
	}
}
