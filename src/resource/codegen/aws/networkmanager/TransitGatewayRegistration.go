package networkmanager

type TransitGatewayRegistration struct {
	// The ARN of the Transit Gateway to register.
	TransitGatewayArn string `json:"transitGatewayArn,omitempty" yaml:"transitGatewayArn,omitempty"`

	// The ID of the Global Network to register to.
	GlobalNetworkId string `json:"globalNetworkId,omitempty" yaml:"globalNetworkId,omitempty"`
}
