package policy

import (
	"math"
	"math/rand"
)

type ActionType int

const (
	IncreaseResolution ActionType = iota
	DecreaseResolution
	IncreaseBitrate
	DecreaseBitrate
	IncreaseFrameRate
	DecreaseFrameRate
)

func (p *Policy) ChooseAction(stateKey string, validActions []ActionType) ActionType {
	p.Mu.RLock()
	defer p.Mu.RUnlock()

	if rand.Float64() < p.Epsilon {
		return validActions[rand.Intn(len(validActions))]
	}

	if qValues, exists := p.QTable[stateKey]; exists {
		var bestAction ActionType
		maxQ := -math.MaxFloat64
		for action, q := range qValues {
			if q > maxQ {
				maxQ = q
				bestAction = action
			}
		}
		return bestAction
	}

	return validActions[rand.Intn(len(validActions))]
}
