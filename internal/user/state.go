package user

import "fmt"

func (u *User) GetStateKey() string {
	return fmt.Sprintf("%s|%t|%s",
		u.NetworkBandwidth,
		u.NetworkDipStatus,
		u.MaxResolution,
	)
}
