package video

import "fmt"

type Video struct {
	Resolution string
	Bitrate    string
	FrameRate  string
}

var ValidResolutions = map[string]bool{
	"720p":  true,
	"1080p": true,
	"4K":    true,
}

var ValidBitrates = map[string]bool{
	"500 Kbps": true,
	"1 Mbps":   true,
	"2.5 Mbps": true,
	"5 Mbps":   true,
}

var ValidFrameRates = map[string]bool{
	"24 fps": true,
	"30 fps": true,
	"60 fps": true,
}

func NewVideo(resolution, bitrate, frameRate string) (*Video, error) {
	if !ValidResolutions[resolution] {
		return nil, fmt.Errorf("invalid resolution: %s. Must be one of: 720p, 1080p, 4K", resolution)
	}

	if !ValidBitrates[bitrate] {
		return nil, fmt.Errorf("invalid bitrate: %s. Must be one of: 500 Kbps, 1 Mbps, 2.5 Mbps, 5 Mbps", bitrate)
	}

	if !ValidFrameRates[frameRate] {
		return nil, fmt.Errorf("invalid frame rate: %s. Must be one of: 24 fps, 30 fps, 60 fps", frameRate)
	}

	return &Video{
		Resolution: resolution,
		Bitrate:    bitrate,
		FrameRate:  frameRate,
	}, nil
}
