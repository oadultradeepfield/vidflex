package simulator

import (
	"log"
	"math/rand"
	"sync"

	"github.com/oadultradeepfield/vidflex/internal/user"
)

type Session struct {
	User              *user.User
	Mu                sync.Mutex
	OriginalBandwidth string
}

func (s *Simulator) AddSessionToSimulator(userID string, user *user.User) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Sessions[userID] = &Session{
		User:              user,
		OriginalBandwidth: user.NetworkBandwidth,
	}
}

func (s *Simulator) SimulateTick(userID string) bool {
	s.Mu.RLock()
	session, exists := s.Sessions[userID]
	s.Mu.RUnlock()

	if !exists {
		log.Printf("Error: Session for userID '%s' does not exist\n", userID)
		return false
	}

	session.Mu.Lock()
	defer session.Mu.Unlock()

	if session.User.NetworkDipStatus {
		if rand.Float64() < s.DipProbability {
			session.User.ToggleNetworkDipstatus()
			session.User.UpdateNetworkBandwidth(session.OriginalBandwidth)
		}
	} else {
		if rand.Float64() < s.DipProbability {
			session.User.ToggleNetworkDipstatus()

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

	return session.User.NetworkDipStatus
}
