package simulator

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/oadultradeepfield/vidflex/internal/user"
)

type Session struct {
	User              *user.User
	DipActive         bool
	DipEndTime        time.Time
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

	now := time.Now()

	if session.DipActive {
		if now.After(session.DipEndTime) {
			session.DipActive = false
			session.User.UpdateNetworkBandwidth(session.OriginalBandwidth)
		}
	}

	if rand.Float64() < s.DipProbability {
		session.DipActive = true
		session.DipEndTime = now.Add(s.DipDuration)

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
