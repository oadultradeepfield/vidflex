package user

import "fmt"

func (u *User) GetStateKey() string {
	return fmt.Sprintf("%s|%t|%s|%s|%s|%s",
		u.NetworkBandwidth,
		u.NetworkDipStatus,
		u.MaxResolution,
		u.CurrentVideo.Resolution,
		u.CurrentVideo.Bitrate,
		u.CurrentVideo.FrameRate,
	)
}
