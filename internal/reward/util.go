package reward

var resolutionBase = map[string]float64{
	"360p":  0.0,
	"720p":  0.33,
	"1080p": 0.66,
	"4K":    1.0,
}

var bitrateBase = map[string]float64{
	"500 Kbps": 500,
	"1 Mbps":   1000,
	"2.5 Mbps": 2500,
	"5 Mbps":   5000,
}

var frameRateBase = map[string]float64{
	"24 fps": 24,
	"30 fps": 30,
	"60 fps": 60,
}

func normalizeResolution(res string) float64 {
	if val, exists := resolutionBase[res]; exists {
		return val
	}
	return 0.0
}

func normalizeBitrate(bitrate string) float64 {
	if val, exists := bitrateBase[bitrate]; exists {
		return val / 5000
	}
	return 0.0
}

func normalizeFrameRate(fps string) float64 {
	if val, exists := frameRateBase[fps]; exists {
		return val / 60
	}
	return 0.0
}
