package types

type Ec2_NetworkInsightsAnalysisExplanationRouteTableRoute struct {
	//
	EgressOnlyInternetGatewayId string `json:"egressOnlyInternetGatewayId,omitempty" yaml:"egressOnlyInternetGatewayId,omitempty"`

	//
	Origin string `json:"origin,omitempty" yaml:"origin,omitempty"`

	//
	GatewayId string `json:"gatewayId,omitempty" yaml:"gatewayId,omitempty"`

	//
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	//
	NatGatewayId string `json:"natGatewayId,omitempty" yaml:"natGatewayId,omitempty"`

	//
	NetworkInterfaceId string `json:"networkInterfaceId,omitempty" yaml:"networkInterfaceId,omitempty"`

	//
	TransitGatewayId string `json:"transitGatewayId,omitempty" yaml:"transitGatewayId,omitempty"`

	//
	VpcPeeringConnectionId string `json:"vpcPeeringConnectionId,omitempty" yaml:"vpcPeeringConnectionId,omitempty"`

	//
	DestinationCidr string `json:"destinationCidr,omitempty" yaml:"destinationCidr,omitempty"`

	//
	DestinationPrefixListId string `json:"destinationPrefixListId,omitempty" yaml:"destinationPrefixListId,omitempty"`
}
