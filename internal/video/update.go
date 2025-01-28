package video

import "fmt"

func (v *Video) UpdateResolution(action string) error {
	return updateField(&v.Resolution, ResolutionOrder, action)
}

func (v *Video) UpdateBitrate(action string) error {
	return updateField(&v.Bitrate, BitrateOrder, action)
}

func (v *Video) UpdateFrameRate(action string) error {
	return updateField(&v.FrameRate, FrameRateOrder, action)
}

func updateField(field *string, order []string, action string) error {
	index := -1
	for i, val := range order {
		if val == *field {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("current value %s is invalid", *field)
	}

	if action == "higher" {
		if index < len(order)-1 {
			*field = order[index+1]
			return nil
		}
		return fmt.Errorf("cannot increase beyond %s", order[len(order)-1])
	} else if action == "lower" {
		if index > 0 {
			*field = order[index-1]
			return nil
		}
		return fmt.Errorf("cannot decrease below %s", order[0])
	}

	return fmt.Errorf("invalid action: %s", action)
}
