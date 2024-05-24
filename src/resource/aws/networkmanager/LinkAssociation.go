package networkmanager

type LinkAssociation struct {
	// The ID of the device.
	DeviceId string `json:"deviceId,omitempty" yaml:"deviceId,omitempty"`

	// The ID of the global network.
	GlobalNetworkId string `json:"globalNetworkId,omitempty" yaml:"globalNetworkId,omitempty"`

	// The ID of the link.
	LinkId string `json:"linkId,omitempty" yaml:"linkId,omitempty"`
}
