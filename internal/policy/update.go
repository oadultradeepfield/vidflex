package policy

func (p *Policy) UpdateQTable(stateKey string, action ActionType, reward float64, nextStateKey string, alpha, gamma float64) {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	if _, exists := p.QTable[stateKey]; !exists {
		p.QTable[stateKey] = make(map[ActionType]float64)
	}

	currentQ := p.QTable[stateKey][action]

	maxNextQ := 0.0
	if nextQValues, exists := p.QTable[nextStateKey]; exists {
		for _, q := range nextQValues {
			if q > maxNextQ {
				maxNextQ = q
			}
		}
	}

	newQ := currentQ + alpha*(reward+gamma*maxNextQ-currentQ)
	p.QTable[stateKey][action] = newQ
}
