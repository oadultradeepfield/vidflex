package user

import (
	"fmt"
)

func (u *User) UpdateNetworkBandwidth(newBandwidth string) error {
	if !ValidNetworkBandwidths[newBandwidth] {
		return fmt.Errorf("invalid network bandwidth: %s. Must be one of: 0-1 Mbps, 1-3 Mbps, 3-5 Mbps, 5+ Mbps", newBandwidth)
	}
	u.NetworkBandwidth = newBandwidth
	return nil
}

func (u *User) UpdateMaxResolution(newDeviceType string) error {
	maxResolution, ok := ValidDeviceTypes[newDeviceType]
	if !ok {
		return fmt.Errorf("invalid device type: %s. Must be one of: Mobile, Tablet, Desktop/TV", newDeviceType)
	}
	u.MaxResolution = maxResolution
	return nil
}
