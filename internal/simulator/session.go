package simulator

import (
	"log"
	"math/rand"
	"sync"

	"github.com/oadultradeepfield/vidflex/internal/user"
)

type Session struct {
	User              *user.User
	DipActive         bool
	DipRemainingTicks int
	Mu                sync.Mutex
	OriginalBandwidth string
}

func (s *Simulator) AddSessionToSimulator(userID string, user *user.User) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Sessions[userID] = &Session{
		User:              user,
		DipActive:         false,
		OriginalBandwidth: user.NetworkBandwidth,
	}
}

func (s *Simulator) SimulateTick(userID string) {
	s.Mu.RLock()
	session, exists := s.Sessions[userID]
	s.Mu.RUnlock()

	if !exists {
		log.Printf("Error: Session for userID '%s' does not exist\n", userID)
		return
	}

	session.Mu.Lock()
	defer session.Mu.Unlock()

	if session.DipActive {
		session.DipRemainingTicks--
		if session.DipRemainingTicks <= 0 {
			session.DipActive = false
			session.User.UpdateNetworkBandwidth(session.OriginalBandwidth)
		}
	}

	if !session.DipActive && rand.Float64() < s.DipProbability {
		session.DipActive = true
		session.DipRemainingTicks = s.DipDurationInTicks

		newBandwidth := "0-1 Mbps"
		switch session.User.NetworkBandwidth {
		case "5+ Mbps":
			newBandwidth = "3-5 Mbps"
		case "3-5 Mbps":
			newBandwidth = "1-3 Mbps"
		case "1-3 Mbps":
			newBandwidth = "0-1 Mbps"
		}

		session.User.UpdateNetworkBandwidth(newBandwidth)
	}
}
