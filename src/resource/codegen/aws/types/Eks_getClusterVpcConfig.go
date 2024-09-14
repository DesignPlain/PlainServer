package types

type Eks_getClusterVpcConfig struct {
	// The cluster security group that was created by Amazon EKS for the cluster.
	ClusterSecurityGroupId string `json:"clusterSecurityGroupId,omitempty" yaml:"clusterSecurityGroupId,omitempty"`

	// Indicates whether or not the Amazon EKS private API server endpoint is enabled.
	EndpointPrivateAccess bool `json:"endpointPrivateAccess,omitempty" yaml:"endpointPrivateAccess,omitempty"`

	// Indicates whether or not the Amazon EKS public API server endpoint is enabled.
	EndpointPublicAccess bool `json:"endpointPublicAccess,omitempty" yaml:"endpointPublicAccess,omitempty"`

	// List of CIDR blocks. Indicates which CIDR blocks can access the Amazon EKS public API server endpoint.
	PublicAccessCidrs []string `json:"publicAccessCidrs,omitempty" yaml:"publicAccessCidrs,omitempty"`

	// List of security group IDs
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// List of subnet IDs
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// The VPC associated with your cluster.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
