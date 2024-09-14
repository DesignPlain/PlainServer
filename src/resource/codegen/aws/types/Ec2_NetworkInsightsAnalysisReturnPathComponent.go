package types

type Ec2_NetworkInsightsAnalysisReturnPathComponent struct {
	//
	InboundHeaders []Ec2_NetworkInsightsAnalysisReturnPathComponentInboundHeader `json:"inboundHeaders,omitempty" yaml:"inboundHeaders,omitempty"`

	//
	OutboundHeaders []Ec2_NetworkInsightsAnalysisReturnPathComponentOutboundHeader `json:"outboundHeaders,omitempty" yaml:"outboundHeaders,omitempty"`

	//
	RouteTableRoutes []Ec2_NetworkInsightsAnalysisReturnPathComponentRouteTableRoute `json:"routeTableRoutes,omitempty" yaml:"routeTableRoutes,omitempty"`

	//
	SecurityGroupRules []Ec2_NetworkInsightsAnalysisReturnPathComponentSecurityGroupRule `json:"securityGroupRules,omitempty" yaml:"securityGroupRules,omitempty"`

	//
	SourceVpcs []Ec2_NetworkInsightsAnalysisReturnPathComponentSourceVpc `json:"sourceVpcs,omitempty" yaml:"sourceVpcs,omitempty"`

	//
	TransitGatewayRouteTableRoutes []Ec2_NetworkInsightsAnalysisReturnPathComponentTransitGatewayRouteTableRoute `json:"transitGatewayRouteTableRoutes,omitempty" yaml:"transitGatewayRouteTableRoutes,omitempty"`

	//
	AdditionalDetails []Ec2_NetworkInsightsAnalysisReturnPathComponentAdditionalDetail `json:"additionalDetails,omitempty" yaml:"additionalDetails,omitempty"`

	//
	DestinationVpcs []Ec2_NetworkInsightsAnalysisReturnPathComponentDestinationVpc `json:"destinationVpcs,omitempty" yaml:"destinationVpcs,omitempty"`

	//
	TransitGateways []Ec2_NetworkInsightsAnalysisReturnPathComponentTransitGateway `json:"transitGateways,omitempty" yaml:"transitGateways,omitempty"`

	//
	Vpcs []Ec2_NetworkInsightsAnalysisReturnPathComponentVpc `json:"vpcs,omitempty" yaml:"vpcs,omitempty"`

	//
	SequenceNumber int `json:"sequenceNumber,omitempty" yaml:"sequenceNumber,omitempty"`

	//
	AttachedTos []Ec2_NetworkInsightsAnalysisReturnPathComponentAttachedTo `json:"attachedTos,omitempty" yaml:"attachedTos,omitempty"`

	//
	Components []Ec2_NetworkInsightsAnalysisReturnPathComponentComponent `json:"components,omitempty" yaml:"components,omitempty"`

	//
	AclRules []Ec2_NetworkInsightsAnalysisReturnPathComponentAclRule `json:"aclRules,omitempty" yaml:"aclRules,omitempty"`

	//
	Subnets []Ec2_NetworkInsightsAnalysisReturnPathComponentSubnet `json:"subnets,omitempty" yaml:"subnets,omitempty"`
}
