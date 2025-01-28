package reward

type Reward struct {
	CollaborativeScores map[string]float64
	CostPenaltyMap      map[string]float64
}

func NewReward() *Reward {
	return &Reward{
		CollaborativeScores: make(map[string]float64),
		CostPenaltyMap: map[string]float64{
			"0-1 Mbps": 0.8,
			"1-3 Mbps": 0.5,
			"3-5 Mbps": 0.3,
			"5+ Mbps":  0.1,
		},
	}
}
