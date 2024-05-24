package types

type Ec2_getNetworkInsightsAnalysisExplanation struct {
	//
	ElasticLoadBalancerListeners []Ec2_getNetworkInsightsAnalysisExplanationElasticLoadBalancerListener `json:"elasticLoadBalancerListeners,omitempty" yaml:"elasticLoadBalancerListeners,omitempty"`

	//
	InternetGateways []Ec2_getNetworkInsightsAnalysisExplanationInternetGateway `json:"internetGateways,omitempty" yaml:"internetGateways,omitempty"`

	//
	PrefixLists []Ec2_getNetworkInsightsAnalysisExplanationPrefixList `json:"prefixLists,omitempty" yaml:"prefixLists,omitempty"`

	//
	Protocols []string `json:"protocols,omitempty" yaml:"protocols,omitempty"`

	//
	RouteTables []Ec2_getNetworkInsightsAnalysisExplanationRouteTable `json:"routeTables,omitempty" yaml:"routeTables,omitempty"`

	//
	SourceVpcs []Ec2_getNetworkInsightsAnalysisExplanationSourceVpc `json:"sourceVpcs,omitempty" yaml:"sourceVpcs,omitempty"`

	//
	TransitGateways []Ec2_getNetworkInsightsAnalysisExplanationTransitGateway `json:"transitGateways,omitempty" yaml:"transitGateways,omitempty"`

	//
	AvailabilityZones []string `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`

	//
	Vpcs []Ec2_getNetworkInsightsAnalysisExplanationVpc `json:"vpcs,omitempty" yaml:"vpcs,omitempty"`

	//
	LoadBalancerListenerPort int `json:"loadBalancerListenerPort,omitempty" yaml:"loadBalancerListenerPort,omitempty"`

	//
	LoadBalancerTargetGroups []Ec2_getNetworkInsightsAnalysisExplanationLoadBalancerTargetGroup `json:"loadBalancerTargetGroups,omitempty" yaml:"loadBalancerTargetGroups,omitempty"`

	//
	SubnetRouteTables []Ec2_getNetworkInsightsAnalysisExplanationSubnetRouteTable `json:"subnetRouteTables,omitempty" yaml:"subnetRouteTables,omitempty"`

	//
	TransitGatewayAttachments []Ec2_getNetworkInsightsAnalysisExplanationTransitGatewayAttachment `json:"transitGatewayAttachments,omitempty" yaml:"transitGatewayAttachments,omitempty"`

	//
	VpnConnections []Ec2_getNetworkInsightsAnalysisExplanationVpnConnection `json:"vpnConnections,omitempty" yaml:"vpnConnections,omitempty"`

	//
	AttachedTos []Ec2_getNetworkInsightsAnalysisExplanationAttachedTo `json:"attachedTos,omitempty" yaml:"attachedTos,omitempty"`

	//
	Destinations []Ec2_getNetworkInsightsAnalysisExplanationDestination `json:"destinations,omitempty" yaml:"destinations,omitempty"`

	//
	LoadBalancerArn string `json:"loadBalancerArn,omitempty" yaml:"loadBalancerArn,omitempty"`

	//
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	//
	SecurityGroupRules []Ec2_getNetworkInsightsAnalysisExplanationSecurityGroupRule `json:"securityGroupRules,omitempty" yaml:"securityGroupRules,omitempty"`

	//
	Acls []Ec2_getNetworkInsightsAnalysisExplanationAcl `json:"acls,omitempty" yaml:"acls,omitempty"`

	//
	Components []Ec2_getNetworkInsightsAnalysisExplanationComponent `json:"components,omitempty" yaml:"components,omitempty"`

	//
	IngressRouteTables []Ec2_getNetworkInsightsAnalysisExplanationIngressRouteTable `json:"ingressRouteTables,omitempty" yaml:"ingressRouteTables,omitempty"`

	//
	LoadBalancerTargetGroup []Ec2_getNetworkInsightsAnalysisExplanationLoadBalancerTargetGroup `json:"loadBalancerTargetGroup,omitempty" yaml:"loadBalancerTargetGroup,omitempty"`

	//
	LoadBalancerTargetPort int `json:"loadBalancerTargetPort,omitempty" yaml:"loadBalancerTargetPort,omitempty"`

	//
	NatGateways []Ec2_getNetworkInsightsAnalysisExplanationNatGateway `json:"natGateways,omitempty" yaml:"natGateways,omitempty"`

	//
	RouteTableRoutes []Ec2_getNetworkInsightsAnalysisExplanationRouteTableRoute `json:"routeTableRoutes,omitempty" yaml:"routeTableRoutes,omitempty"`

	//
	TransitGatewayRouteTables []Ec2_getNetworkInsightsAnalysisExplanationTransitGatewayRouteTable `json:"transitGatewayRouteTables,omitempty" yaml:"transitGatewayRouteTables,omitempty"`

	//
	ClassicLoadBalancerListeners []Ec2_getNetworkInsightsAnalysisExplanationClassicLoadBalancerListener `json:"classicLoadBalancerListeners,omitempty" yaml:"classicLoadBalancerListeners,omitempty"`

	//
	DestinationVpcs []Ec2_getNetworkInsightsAnalysisExplanationDestinationVpc `json:"destinationVpcs,omitempty" yaml:"destinationVpcs,omitempty"`

	//
	ExplanationCode string `json:"explanationCode,omitempty" yaml:"explanationCode,omitempty"`

	//
	PacketField string `json:"packetField,omitempty" yaml:"packetField,omitempty"`

	//
	SecurityGroups []Ec2_getNetworkInsightsAnalysisExplanationSecurityGroup `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	//
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	//
	TransitGatewayRouteTableRoutes []Ec2_getNetworkInsightsAnalysisExplanationTransitGatewayRouteTableRoute `json:"transitGatewayRouteTableRoutes,omitempty" yaml:"transitGatewayRouteTableRoutes,omitempty"`

	//
	CustomerGateways []Ec2_getNetworkInsightsAnalysisExplanationCustomerGateway `json:"customerGateways,omitempty" yaml:"customerGateways,omitempty"`

	//
	Cidrs []string `json:"cidrs,omitempty" yaml:"cidrs,omitempty"`

	//
	Direction string `json:"direction,omitempty" yaml:"direction,omitempty"`

	//
	Subnets []Ec2_getNetworkInsightsAnalysisExplanationSubnet `json:"subnets,omitempty" yaml:"subnets,omitempty"`

	//
	VpnGateways []Ec2_getNetworkInsightsAnalysisExplanationVpnGateway `json:"vpnGateways,omitempty" yaml:"vpnGateways,omitempty"`

	//
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	//
	SecurityGroup []Ec2_getNetworkInsightsAnalysisExplanationSecurityGroup `json:"securityGroup,omitempty" yaml:"securityGroup,omitempty"`

	//
	VpcEndpoints []Ec2_getNetworkInsightsAnalysisExplanationVpcEndpoint `json:"vpcEndpoints,omitempty" yaml:"vpcEndpoints,omitempty"`

	//
	NetworkInterfaces []Ec2_getNetworkInsightsAnalysisExplanationNetworkInterface `json:"networkInterfaces,omitempty" yaml:"networkInterfaces,omitempty"`

	//
	Addresses []string `json:"addresses,omitempty" yaml:"addresses,omitempty"`

	//
	MissingComponent string `json:"missingComponent,omitempty" yaml:"missingComponent,omitempty"`

	//
	PortRanges []Ec2_getNetworkInsightsAnalysisExplanationPortRange `json:"portRanges,omitempty" yaml:"portRanges,omitempty"`

	//
	VpcPeeringConnections []Ec2_getNetworkInsightsAnalysisExplanationVpcPeeringConnection `json:"vpcPeeringConnections,omitempty" yaml:"vpcPeeringConnections,omitempty"`

	//
	AclRules []Ec2_getNetworkInsightsAnalysisExplanationAclRule `json:"aclRules,omitempty" yaml:"aclRules,omitempty"`
}
