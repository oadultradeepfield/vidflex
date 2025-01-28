package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/oadultradeepfield/vidflex/internal/policy"
	"github.com/oadultradeepfield/vidflex/internal/reward"
	"github.com/oadultradeepfield/vidflex/internal/simulator"
	"github.com/oadultradeepfield/vidflex/internal/user"
)

func main() {
	var (
		episodes        = flag.Int("episodes", 1000, "Number of training episodes")
		usersPerEpisode = flag.Int("users", 1000, "Users per episode")
		alpha           = flag.Float64("alpha", 0.1, "Learning rate (0.0-1.0)")
		gamma           = flag.Float64("gamma", 0.9, "Discount factor (0.0-1.0)")
		epsilon         = flag.Float64("epsilon", 0.1, "Initial exploration rate")
		decayEvery      = flag.Int("decay-every", 100, "Epsilon decay interval")
		decayRate       = flag.Float64("decay-rate", 0.1, "Epsilon decay rate")
	)
	flag.Parse()

	if *alpha < 0 || *alpha > 1 || *gamma < 0 || *gamma > 1 || *epsilon < 0 || *epsilon > 1 {
		log.Fatal("Parameters must be between 0 and 1")
	}

	p := policy.NewPolicy()
	p.Epsilon = *epsilon
	p.DecayEvery = *decayEvery
	p.DecayRate = *decayRate

	r := reward.NewReward()
	s := simulator.NewSimulator()

	start := time.Now()
	for ep := 0; ep < *episodes; ep++ {
		var wg sync.WaitGroup
		results := make(chan float64, *usersPerEpisode)
		errors := make(chan error, *usersPerEpisode)

		for i := 0; i < *usersPerEpisode; i++ {
			wg.Add(1)
			go func(userID string) {
				defer wg.Done()

				u, action, nextStateKey, err := simulateUserSession(userID, p, s)
				if err != nil {
					errors <- fmt.Errorf("session failed for %s: %w", userID, err)
					return
				}

				s.SimulateTick(userID)
				reward := r.CalculateReward(u)
				results <- reward

				p.UpdateQTable(
					u.GetStateKey(),
					action,
					reward,
					nextStateKey,
					*alpha,
					*gamma,
				)
			}(fmt.Sprintf("user-%d-%d", ep, i))
		}

		wg.Wait()
		close(results)
		close(errors)

		for err := range errors {
			log.Println(err)
		}

		avgReward := calculateAverage(results)

		if ep > 0 && ep%p.DecayEvery == 0 {
			p.Epsilon *= (1 - p.DecayRate)
		}

		if ep%100 == 0 {
			fmt.Printf("Episode %d | Average Reward: %.4f | Epsilon: %.4f\n",
				ep, avgReward, p.Epsilon)
		}
	}

	fmt.Printf("Training completed in %v\n", time.Since(start))
}

func simulateUserSession(userID string, p *policy.Policy, simulator *simulator.Simulator) (*user.User, policy.ActionType, string, error) {
	bandwidths := []string{"0-1 Mbps", "1-3 Mbps", "3-5 Mbps", "5+ Mbps"}
	deviceTypes := []string{"Mobile", "Tablet", "Desktop/TV"}

	networkBandwidth := bandwidths[rand.Intn(len(bandwidths))]
	deviceType := deviceTypes[rand.Intn(len(deviceTypes))]

	u, err := user.NewUser(networkBandwidth, deviceType)
	if err != nil {
		return nil, policy.DecreaseResolution, "", fmt.Errorf("failed to create user: %w", err)
	}

	simulator.AddSessionToSimulator(userID, u)

	stateKey := u.GetStateKey()
	validActions := []policy.ActionType{
		policy.IncreaseResolution,
		policy.DecreaseResolution,
		policy.IncreaseBitrate,
		policy.DecreaseBitrate,
		policy.IncreaseFrameRate,
		policy.DecreaseFrameRate,
	}
	action := p.ChooseAction(stateKey, validActions)

	applyAction(u, action)
	nextStateKey := u.GetStateKey()

	return u, action, nextStateKey, nil
}

func applyAction(u *user.User, action policy.ActionType) {
	switch action {
	case policy.IncreaseResolution:
		_ = u.CurrentVideo.UpdateResolution("higher")
	case policy.DecreaseResolution:
		_ = u.CurrentVideo.UpdateResolution("lower")
	case policy.IncreaseBitrate:
		_ = u.CurrentVideo.UpdateBitrate("higher")
	case policy.DecreaseBitrate:
		_ = u.CurrentVideo.UpdateBitrate("lower")
	case policy.IncreaseFrameRate:
		_ = u.CurrentVideo.UpdateFrameRate("higher")
	case policy.DecreaseFrameRate:
		_ = u.CurrentVideo.UpdateFrameRate("lower")
	}
}

func calculateAverage(results <-chan float64) float64 {
	var sum float64
	var count int

	for r := range results {
		sum += r
		count++
	}

	if count == 0 {
		return 0
	}
	return sum / float64(count)
}
