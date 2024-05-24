package redshift

type EndpointAccess struct {
	// The cluster identifier of the cluster to access.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`

	// The Redshift-managed VPC endpoint name.
	EndpointName string `json:"endpointName,omitempty" yaml:"endpointName,omitempty"`

	// The Amazon Web Services account ID of the owner of the cluster. This is only required if the cluster is in another Amazon Web Services account.
	ResourceOwner string `json:"resourceOwner,omitempty" yaml:"resourceOwner,omitempty"`

	// The subnet group from which Amazon Redshift chooses the subnet to deploy the endpoint.
	SubnetGroupName string `json:"subnetGroupName,omitempty" yaml:"subnetGroupName,omitempty"`

	// The security group that defines the ports, protocols, and sources for inbound traffic that you are authorizing into your endpoint.
	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" yaml:"vpcSecurityGroupIds,omitempty"`
}
