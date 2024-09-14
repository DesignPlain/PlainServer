package types

type Finspace_KxEnvironmentTransitGatewayConfiguration struct {
	// Rules that define how you manage outbound traffic from kdb network to your internal network. Defined below.
	AttachmentNetworkAclConfigurations []Finspace_KxEnvironmentTransitGatewayConfigurationAttachmentNetworkAclConfiguration `json:"attachmentNetworkAclConfigurations,omitempty" yaml:"attachmentNetworkAclConfigurations,omitempty"`

	// Routing CIDR on behalf of KX environment. It could be any “/26 range in the 100.64.0.0 CIDR space. After providing, it will be added to the customer’s transit gateway routing table so that the traffics could be routed to KX network.
	RoutableCidrSpace string `json:"routableCidrSpace,omitempty" yaml:"routableCidrSpace,omitempty"`

	// Identifier of the transit gateway created by the customer to connect outbound traffics from KX network to your internal network.
	TransitGatewayId string `json:"transitGatewayId,omitempty" yaml:"transitGatewayId,omitempty"`
}
