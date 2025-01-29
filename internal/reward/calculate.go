package reward

import (
	"github.com/oadultradeepfield/vidflex/internal/user"
)

const (
	resolutionWeight = 0.4
	bitrateWeight    = 0.3
	frameRateWeight  = 0.3
)

func (r *Reward) CalculateReward(u *user.User, dipStatus bool) (float64, float64) {
	engagement := r.calculateEngagement(u)
	if dipStatus {
		engagement *= 0.7
	}

	costPenalty := r.CostPenaltyMap[u.CurrentVideo.Bitrate]
	collaborativeFilteringBonus := 0.2 * r.CollaborativeScores[u.GetStateKey()]
	return engagement + collaborativeFilteringBonus, costPenalty
}

func (r *Reward) calculateEngagement(u *user.User) float64 {
	normalizedResolution := normalizeResolution(u.CurrentVideo.Resolution)
	normalizedBitrate := normalizeBitrate(u.CurrentVideo.Bitrate)
	normalizedFrameRate := normalizeFrameRate(u.CurrentVideo.FrameRate)
	return normalizedResolution*resolutionWeight + normalizedBitrate*bitrateWeight + normalizedFrameRate*frameRateWeight
}
