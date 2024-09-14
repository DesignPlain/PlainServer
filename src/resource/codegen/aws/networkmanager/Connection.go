package networkmanager

type Connection struct {
	// A description of the connection.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The ID of the first device in the connection.
	DeviceId string `json:"deviceId,omitempty" yaml:"deviceId,omitempty"`

	// The ID of the global network.
	GlobalNetworkId string `json:"globalNetworkId,omitempty" yaml:"globalNetworkId,omitempty"`

	// The ID of the link for the first device.
	LinkId string `json:"linkId,omitempty" yaml:"linkId,omitempty"`

	// Key-value tags for the connection. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ID of the second device in the connection.
	ConnectedDeviceId string `json:"connectedDeviceId,omitempty" yaml:"connectedDeviceId,omitempty"`

	// The ID of the link for the second device.
	ConnectedLinkId string `json:"connectedLinkId,omitempty" yaml:"connectedLinkId,omitempty"`
}
