package policy

import (
	"sync"
)

type Policy struct {
	Epsilon    float64
	QTable     map[string]map[ActionType]float64
	DecayRate  float64
	DecayEvery int
	Mu         sync.RWMutex
}

func NewPolicy() *Policy {
	return &Policy{
		Epsilon:    0.8,
		QTable:     make(map[string]map[ActionType]float64),
		DecayRate:  0.05,
		DecayEvery: 100,
	}
}
