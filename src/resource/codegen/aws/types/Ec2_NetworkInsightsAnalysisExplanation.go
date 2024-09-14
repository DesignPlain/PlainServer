package types

type Ec2_NetworkInsightsAnalysisExplanation struct {
	//
	LoadBalancerTargetGroup []Ec2_NetworkInsightsAnalysisExplanationLoadBalancerTargetGroup `json:"loadBalancerTargetGroup,omitempty" yaml:"loadBalancerTargetGroup,omitempty"`

	//
	LoadBalancerTargetPort int `json:"loadBalancerTargetPort,omitempty" yaml:"loadBalancerTargetPort,omitempty"`

	//
	SecurityGroupRules []Ec2_NetworkInsightsAnalysisExplanationSecurityGroupRule `json:"securityGroupRules,omitempty" yaml:"securityGroupRules,omitempty"`

	//
	SourceVpcs []Ec2_NetworkInsightsAnalysisExplanationSourceVpc `json:"sourceVpcs,omitempty" yaml:"sourceVpcs,omitempty"`

	//
	Acls []Ec2_NetworkInsightsAnalysisExplanationAcl `json:"acls,omitempty" yaml:"acls,omitempty"`

	//
	Addresses []string `json:"addresses,omitempty" yaml:"addresses,omitempty"`

	//
	CustomerGateways []Ec2_NetworkInsightsAnalysisExplanationCustomerGateway `json:"customerGateways,omitempty" yaml:"customerGateways,omitempty"`

	//
	InternetGateways []Ec2_NetworkInsightsAnalysisExplanationInternetGateway `json:"internetGateways,omitempty" yaml:"internetGateways,omitempty"`

	//
	PrefixLists []Ec2_NetworkInsightsAnalysisExplanationPrefixList `json:"prefixLists,omitempty" yaml:"prefixLists,omitempty"`

	//
	ClassicLoadBalancerListeners []Ec2_NetworkInsightsAnalysisExplanationClassicLoadBalancerListener `json:"classicLoadBalancerListeners,omitempty" yaml:"classicLoadBalancerListeners,omitempty"`

	//
	ExplanationCode string `json:"explanationCode,omitempty" yaml:"explanationCode,omitempty"`

	//
	IngressRouteTables []Ec2_NetworkInsightsAnalysisExplanationIngressRouteTable `json:"ingressRouteTables,omitempty" yaml:"ingressRouteTables,omitempty"`

	//
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	//
	RouteTableRoutes []Ec2_NetworkInsightsAnalysisExplanationRouteTableRoute `json:"routeTableRoutes,omitempty" yaml:"routeTableRoutes,omitempty"`

	//
	RouteTables []Ec2_NetworkInsightsAnalysisExplanationRouteTable `json:"routeTables,omitempty" yaml:"routeTables,omitempty"`

	//
	AclRules []Ec2_NetworkInsightsAnalysisExplanationAclRule `json:"aclRules,omitempty" yaml:"aclRules,omitempty"`

	//
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	//
	DestinationVpcs []Ec2_NetworkInsightsAnalysisExplanationDestinationVpc `json:"destinationVpcs,omitempty" yaml:"destinationVpcs,omitempty"`

	//
	NatGateways []Ec2_NetworkInsightsAnalysisExplanationNatGateway `json:"natGateways,omitempty" yaml:"natGateways,omitempty"`

	//
	NetworkInterfaces []Ec2_NetworkInsightsAnalysisExplanationNetworkInterface `json:"networkInterfaces,omitempty" yaml:"networkInterfaces,omitempty"`

	//
	VpcEndpoints []Ec2_NetworkInsightsAnalysisExplanationVpcEndpoint `json:"vpcEndpoints,omitempty" yaml:"vpcEndpoints,omitempty"`

	//
	VpnConnections []Ec2_NetworkInsightsAnalysisExplanationVpnConnection `json:"vpnConnections,omitempty" yaml:"vpnConnections,omitempty"`

	//
	AttachedTos []Ec2_NetworkInsightsAnalysisExplanationAttachedTo `json:"attachedTos,omitempty" yaml:"attachedTos,omitempty"`

	//
	ElasticLoadBalancerListeners []Ec2_NetworkInsightsAnalysisExplanationElasticLoadBalancerListener `json:"elasticLoadBalancerListeners,omitempty" yaml:"elasticLoadBalancerListeners,omitempty"`

	//
	LoadBalancerTargetGroups []Ec2_NetworkInsightsAnalysisExplanationLoadBalancerTargetGroup `json:"loadBalancerTargetGroups,omitempty" yaml:"loadBalancerTargetGroups,omitempty"`

	//
	MissingComponent string `json:"missingComponent,omitempty" yaml:"missingComponent,omitempty"`

	//
	TransitGatewayAttachments []Ec2_NetworkInsightsAnalysisExplanationTransitGatewayAttachment `json:"transitGatewayAttachments,omitempty" yaml:"transitGatewayAttachments,omitempty"`

	//
	TransitGatewayRouteTableRoutes []Ec2_NetworkInsightsAnalysisExplanationTransitGatewayRouteTableRoute `json:"transitGatewayRouteTableRoutes,omitempty" yaml:"transitGatewayRouteTableRoutes,omitempty"`

	//
	TransitGateways []Ec2_NetworkInsightsAnalysisExplanationTransitGateway `json:"transitGateways,omitempty" yaml:"transitGateways,omitempty"`

	//
	Cidrs []string `json:"cidrs,omitempty" yaml:"cidrs,omitempty"`

	//
	Components []Ec2_NetworkInsightsAnalysisExplanationComponent `json:"components,omitempty" yaml:"components,omitempty"`

	//
	Destinations []Ec2_NetworkInsightsAnalysisExplanationDestination `json:"destinations,omitempty" yaml:"destinations,omitempty"`

	//
	Direction string `json:"direction,omitempty" yaml:"direction,omitempty"`

	//
	SecurityGroup []Ec2_NetworkInsightsAnalysisExplanationSecurityGroup `json:"securityGroup,omitempty" yaml:"securityGroup,omitempty"`

	//
	SecurityGroups []Ec2_NetworkInsightsAnalysisExplanationSecurityGroup `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	//
	Vpcs []Ec2_NetworkInsightsAnalysisExplanationVpc `json:"vpcs,omitempty" yaml:"vpcs,omitempty"`

	//
	VpnGateways []Ec2_NetworkInsightsAnalysisExplanationVpnGateway `json:"vpnGateways,omitempty" yaml:"vpnGateways,omitempty"`

	//
	AvailabilityZones []string `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`

	//
	SubnetRouteTables []Ec2_NetworkInsightsAnalysisExplanationSubnetRouteTable `json:"subnetRouteTables,omitempty" yaml:"subnetRouteTables,omitempty"`

	//
	TransitGatewayRouteTables []Ec2_NetworkInsightsAnalysisExplanationTransitGatewayRouteTable `json:"transitGatewayRouteTables,omitempty" yaml:"transitGatewayRouteTables,omitempty"`

	//
	VpcPeeringConnections []Ec2_NetworkInsightsAnalysisExplanationVpcPeeringConnection `json:"vpcPeeringConnections,omitempty" yaml:"vpcPeeringConnections,omitempty"`

	//
	LoadBalancerArn string `json:"loadBalancerArn,omitempty" yaml:"loadBalancerArn,omitempty"`

	//
	LoadBalancerListenerPort int `json:"loadBalancerListenerPort,omitempty" yaml:"loadBalancerListenerPort,omitempty"`

	//
	PacketField string `json:"packetField,omitempty" yaml:"packetField,omitempty"`

	//
	PortRanges []Ec2_NetworkInsightsAnalysisExplanationPortRange `json:"portRanges,omitempty" yaml:"portRanges,omitempty"`

	//
	Protocols []string `json:"protocols,omitempty" yaml:"protocols,omitempty"`

	//
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	//
	Subnets []Ec2_NetworkInsightsAnalysisExplanationSubnet `json:"subnets,omitempty" yaml:"subnets,omitempty"`
}
