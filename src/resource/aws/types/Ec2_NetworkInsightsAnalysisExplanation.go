package types

type Ec2_NetworkInsightsAnalysisExplanation struct {
	//
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	//
	SecurityGroupRules []Ec2_NetworkInsightsAnalysisExplanationSecurityGroupRule `json:"securityGroupRules,omitempty" yaml:"securityGroupRules,omitempty"`

	//
	Components []Ec2_NetworkInsightsAnalysisExplanationComponent `json:"components,omitempty" yaml:"components,omitempty"`

	//
	CustomerGateways []Ec2_NetworkInsightsAnalysisExplanationCustomerGateway `json:"customerGateways,omitempty" yaml:"customerGateways,omitempty"`

	//
	DestinationVpcs []Ec2_NetworkInsightsAnalysisExplanationDestinationVpc `json:"destinationVpcs,omitempty" yaml:"destinationVpcs,omitempty"`

	//
	Direction string `json:"direction,omitempty" yaml:"direction,omitempty"`

	//
	LoadBalancerArn string `json:"loadBalancerArn,omitempty" yaml:"loadBalancerArn,omitempty"`

	//
	NatGateways []Ec2_NetworkInsightsAnalysisExplanationNatGateway `json:"natGateways,omitempty" yaml:"natGateways,omitempty"`

	//
	Subnets []Ec2_NetworkInsightsAnalysisExplanationSubnet `json:"subnets,omitempty" yaml:"subnets,omitempty"`

	//
	Acls []Ec2_NetworkInsightsAnalysisExplanationAcl `json:"acls,omitempty" yaml:"acls,omitempty"`

	//
	Cidrs []string `json:"cidrs,omitempty" yaml:"cidrs,omitempty"`

	//
	IngressRouteTables []Ec2_NetworkInsightsAnalysisExplanationIngressRouteTable `json:"ingressRouteTables,omitempty" yaml:"ingressRouteTables,omitempty"`

	//
	Protocols []string `json:"protocols,omitempty" yaml:"protocols,omitempty"`

	//
	SubnetRouteTables []Ec2_NetworkInsightsAnalysisExplanationSubnetRouteTable `json:"subnetRouteTables,omitempty" yaml:"subnetRouteTables,omitempty"`

	//
	AttachedTos []Ec2_NetworkInsightsAnalysisExplanationAttachedTo `json:"attachedTos,omitempty" yaml:"attachedTos,omitempty"`

	//
	InternetGateways []Ec2_NetworkInsightsAnalysisExplanationInternetGateway `json:"internetGateways,omitempty" yaml:"internetGateways,omitempty"`

	//
	SourceVpcs []Ec2_NetworkInsightsAnalysisExplanationSourceVpc `json:"sourceVpcs,omitempty" yaml:"sourceVpcs,omitempty"`

	//
	Addresses []string `json:"addresses,omitempty" yaml:"addresses,omitempty"`

	//
	ExplanationCode string `json:"explanationCode,omitempty" yaml:"explanationCode,omitempty"`

	//
	LoadBalancerTargetGroup []Ec2_NetworkInsightsAnalysisExplanationLoadBalancerTargetGroup `json:"loadBalancerTargetGroup,omitempty" yaml:"loadBalancerTargetGroup,omitempty"`

	//
	SecurityGroups []Ec2_NetworkInsightsAnalysisExplanationSecurityGroup `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	//
	VpcEndpoints []Ec2_NetworkInsightsAnalysisExplanationVpcEndpoint `json:"vpcEndpoints,omitempty" yaml:"vpcEndpoints,omitempty"`

	//
	AclRules []Ec2_NetworkInsightsAnalysisExplanationAclRule `json:"aclRules,omitempty" yaml:"aclRules,omitempty"`

	//
	ElasticLoadBalancerListeners []Ec2_NetworkInsightsAnalysisExplanationElasticLoadBalancerListener `json:"elasticLoadBalancerListeners,omitempty" yaml:"elasticLoadBalancerListeners,omitempty"`

	//
	LoadBalancerTargetGroups []Ec2_NetworkInsightsAnalysisExplanationLoadBalancerTargetGroup `json:"loadBalancerTargetGroups,omitempty" yaml:"loadBalancerTargetGroups,omitempty"`

	//
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	//
	TransitGatewayRouteTableRoutes []Ec2_NetworkInsightsAnalysisExplanationTransitGatewayRouteTableRoute `json:"transitGatewayRouteTableRoutes,omitempty" yaml:"transitGatewayRouteTableRoutes,omitempty"`

	//
	TransitGatewayRouteTables []Ec2_NetworkInsightsAnalysisExplanationTransitGatewayRouteTable `json:"transitGatewayRouteTables,omitempty" yaml:"transitGatewayRouteTables,omitempty"`

	//
	RouteTables []Ec2_NetworkInsightsAnalysisExplanationRouteTable `json:"routeTables,omitempty" yaml:"routeTables,omitempty"`

	//
	SecurityGroup []Ec2_NetworkInsightsAnalysisExplanationSecurityGroup `json:"securityGroup,omitempty" yaml:"securityGroup,omitempty"`

	//
	Destinations []Ec2_NetworkInsightsAnalysisExplanationDestination `json:"destinations,omitempty" yaml:"destinations,omitempty"`

	//
	LoadBalancerTargetPort int `json:"loadBalancerTargetPort,omitempty" yaml:"loadBalancerTargetPort,omitempty"`

	//
	PacketField string `json:"packetField,omitempty" yaml:"packetField,omitempty"`

	//
	PortRanges []Ec2_NetworkInsightsAnalysisExplanationPortRange `json:"portRanges,omitempty" yaml:"portRanges,omitempty"`

	//
	PrefixLists []Ec2_NetworkInsightsAnalysisExplanationPrefixList `json:"prefixLists,omitempty" yaml:"prefixLists,omitempty"`

	//
	RouteTableRoutes []Ec2_NetworkInsightsAnalysisExplanationRouteTableRoute `json:"routeTableRoutes,omitempty" yaml:"routeTableRoutes,omitempty"`

	//
	TransitGatewayAttachments []Ec2_NetworkInsightsAnalysisExplanationTransitGatewayAttachment `json:"transitGatewayAttachments,omitempty" yaml:"transitGatewayAttachments,omitempty"`

	//
	TransitGateways []Ec2_NetworkInsightsAnalysisExplanationTransitGateway `json:"transitGateways,omitempty" yaml:"transitGateways,omitempty"`

	//
	VpnConnections []Ec2_NetworkInsightsAnalysisExplanationVpnConnection `json:"vpnConnections,omitempty" yaml:"vpnConnections,omitempty"`

	//
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	//
	ClassicLoadBalancerListeners []Ec2_NetworkInsightsAnalysisExplanationClassicLoadBalancerListener `json:"classicLoadBalancerListeners,omitempty" yaml:"classicLoadBalancerListeners,omitempty"`

	//
	NetworkInterfaces []Ec2_NetworkInsightsAnalysisExplanationNetworkInterface `json:"networkInterfaces,omitempty" yaml:"networkInterfaces,omitempty"`

	//
	Vpcs []Ec2_NetworkInsightsAnalysisExplanationVpc `json:"vpcs,omitempty" yaml:"vpcs,omitempty"`

	//
	VpnGateways []Ec2_NetworkInsightsAnalysisExplanationVpnGateway `json:"vpnGateways,omitempty" yaml:"vpnGateways,omitempty"`

	//
	AvailabilityZones []string `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`

	//
	LoadBalancerListenerPort int `json:"loadBalancerListenerPort,omitempty" yaml:"loadBalancerListenerPort,omitempty"`

	//
	MissingComponent string `json:"missingComponent,omitempty" yaml:"missingComponent,omitempty"`

	//
	VpcPeeringConnections []Ec2_NetworkInsightsAnalysisExplanationVpcPeeringConnection `json:"vpcPeeringConnections,omitempty" yaml:"vpcPeeringConnections,omitempty"`
}
