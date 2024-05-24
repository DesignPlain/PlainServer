package networkmanager

type CustomerGatewayAssociation struct {
	// The ID of the global network.
	GlobalNetworkId string `json:"globalNetworkId,omitempty" yaml:"globalNetworkId,omitempty"`

	// The ID of the link.
	LinkId string `json:"linkId,omitempty" yaml:"linkId,omitempty"`

	// The Amazon Resource Name (ARN) of the customer gateway.
	CustomerGatewayArn string `json:"customerGatewayArn,omitempty" yaml:"customerGatewayArn,omitempty"`

	// The ID of the device.
	DeviceId string `json:"deviceId,omitempty" yaml:"deviceId,omitempty"`
}
