package types

type Ec2_getNetworkInsightsAnalysisForwardPathComponent struct {
	//
	Components []Ec2_getNetworkInsightsAnalysisForwardPathComponentComponent `json:"components,omitempty" yaml:"components,omitempty"`

	//
	OutboundHeaders []Ec2_getNetworkInsightsAnalysisForwardPathComponentOutboundHeader `json:"outboundHeaders,omitempty" yaml:"outboundHeaders,omitempty"`

	//
	SecurityGroupRules []Ec2_getNetworkInsightsAnalysisForwardPathComponentSecurityGroupRule `json:"securityGroupRules,omitempty" yaml:"securityGroupRules,omitempty"`

	//
	RouteTableRoutes []Ec2_getNetworkInsightsAnalysisForwardPathComponentRouteTableRoute `json:"routeTableRoutes,omitempty" yaml:"routeTableRoutes,omitempty"`

	//
	TransitGatewayRouteTableRoutes []Ec2_getNetworkInsightsAnalysisForwardPathComponentTransitGatewayRouteTableRoute `json:"transitGatewayRouteTableRoutes,omitempty" yaml:"transitGatewayRouteTableRoutes,omitempty"`

	//
	Vpcs []Ec2_getNetworkInsightsAnalysisForwardPathComponentVpc `json:"vpcs,omitempty" yaml:"vpcs,omitempty"`

	//
	AdditionalDetails []Ec2_getNetworkInsightsAnalysisForwardPathComponentAdditionalDetail `json:"additionalDetails,omitempty" yaml:"additionalDetails,omitempty"`

	//
	InboundHeaders []Ec2_getNetworkInsightsAnalysisForwardPathComponentInboundHeader `json:"inboundHeaders,omitempty" yaml:"inboundHeaders,omitempty"`

	//
	SequenceNumber int `json:"sequenceNumber,omitempty" yaml:"sequenceNumber,omitempty"`

	//
	SourceVpcs []Ec2_getNetworkInsightsAnalysisForwardPathComponentSourceVpc `json:"sourceVpcs,omitempty" yaml:"sourceVpcs,omitempty"`

	//
	TransitGateways []Ec2_getNetworkInsightsAnalysisForwardPathComponentTransitGateway `json:"transitGateways,omitempty" yaml:"transitGateways,omitempty"`

	//
	AclRules []Ec2_getNetworkInsightsAnalysisForwardPathComponentAclRule `json:"aclRules,omitempty" yaml:"aclRules,omitempty"`

	//
	AttachedTos []Ec2_getNetworkInsightsAnalysisForwardPathComponentAttachedTo `json:"attachedTos,omitempty" yaml:"attachedTos,omitempty"`

	//
	DestinationVpcs []Ec2_getNetworkInsightsAnalysisForwardPathComponentDestinationVpc `json:"destinationVpcs,omitempty" yaml:"destinationVpcs,omitempty"`

	//
	Subnets []Ec2_getNetworkInsightsAnalysisForwardPathComponentSubnet `json:"subnets,omitempty" yaml:"subnets,omitempty"`
}
