package user

import (
	"fmt"

	"github.com/oadultradeepfield/vidflex/internal/video"
)

type User struct {
	NetworkBandwidth string
	MaxResolution    string
	CurrentVideo     *video.Video
}

var ValidNetworkBandwidths = map[string]bool{
	"0-1 Mbps": true,
	"1-3 Mbps": true,
	"3-5 Mbps": true,
	"5+ Mbps":  true,
}

var ValidDeviceTypes = map[string]string{
	"Mobile":     "720p",
	"Tablet":     "1080p",
	"Desktop/TV": "4K",
}

func NewUser(networkBandwidth, deviceType string) (*User, error) {
	if !ValidNetworkBandwidths[networkBandwidth] {
		return nil, fmt.Errorf("invalid network bandwidth: %s. Must be one of: 0-1 Mbps, 1-3 Mbps, 3-5 Mbps, 5+ Mbps", networkBandwidth)
	}

	maxResolution, ok := ValidDeviceTypes[deviceType]
	if !ok {
		return nil, fmt.Errorf("invalid device type: %s. Must be one of: Mobile, Tablet, Desktop/TV", deviceType)
	}

	defaultVideo, err := video.NewVideo("720p", "1 Mbps", "30 fps")
	if err != nil {
		return nil, fmt.Errorf("failed to create default video parameters: %w", err)
	}

	return &User{
		NetworkBandwidth: networkBandwidth,
		MaxResolution:    maxResolution,
		CurrentVideo:     defaultVideo,
	}, nil
}
