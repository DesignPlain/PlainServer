package types

type Eks_ClusterVpcConfig struct {
	// ID of the VPC associated with your cluster.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// Cluster security group that was created by Amazon EKS for the cluster. Managed node groups use this security group for control-plane-to-data-plane communication.
	ClusterSecurityGroupId string `json:"clusterSecurityGroupId,omitempty" yaml:"clusterSecurityGroupId,omitempty"`

	// Whether the Amazon EKS private API server endpoint is enabled. Default is `false`.
	EndpointPrivateAccess bool `json:"endpointPrivateAccess,omitempty" yaml:"endpointPrivateAccess,omitempty"`

	// Whether the Amazon EKS public API server endpoint is enabled. Default is `true`.
	EndpointPublicAccess bool `json:"endpointPublicAccess,omitempty" yaml:"endpointPublicAccess,omitempty"`

	// List of CIDR blocks. Indicates which CIDR blocks can access the Amazon EKS public API server endpoint when enabled. EKS defaults this to a list with `0.0.0.0/0`. The provider will only perform drift detection of its value when present in a configuration.
	PublicAccessCidrs []string `json:"publicAccessCidrs,omitempty" yaml:"publicAccessCidrs,omitempty"`

	// List of security group IDs for the cross-account elastic network interfaces that Amazon EKS creates to use to allow communication between your worker nodes and the Kubernetes control plane.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// List of subnet IDs. Must be in at least two different availability zones. Amazon EKS creates cross-account elastic network interfaces in these subnets to allow communication between your worker nodes and the Kubernetes control plane.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`
}
