package video

import "fmt"

func (v *Video) UpdateResolution(newResolution string) error {
	if !ValidResolutions[newResolution] {
		return fmt.Errorf("invalid resolution: %s. Must be one of: 720p, 1080p, 4K", newResolution)
	}
	v.Resolution = newResolution
	return nil
}

func (v *Video) UpdateBitrate(newBitrate string) error {
	if !ValidBitrates[newBitrate] {
		return fmt.Errorf("invalid bitrate: %s. Must be one of: 500 Kbps, 1 Mbps, 2.5 Mbps, 5 Mbps", newBitrate)
	}
	v.Bitrate = newBitrate
	return nil
}

func (v *Video) UpdateFrameRate(newFrameRate string) error {
	if !ValidFrameRates[newFrameRate] {
		return fmt.Errorf("invalid frame rate: %s. Must be one of: 24 fps, 30 fps, 60 fps", newFrameRate)
	}
	v.FrameRate = newFrameRate
	return nil
}
