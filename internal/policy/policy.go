package policy

type Policy struct {
	Epsilon      float64
	QTable       map[string]map[ActionType]float64
	DecayRate    float64
	DecayEvery   int
	EpisodeCount int
}

func NewPolicy(initialEpsilon float64) *Policy {
	return &Policy{
		Epsilon:    initialEpsilon,
		QTable:     make(map[string]map[ActionType]float64),
		DecayRate:  0.05,
		DecayEvery: 100,
	}
}
