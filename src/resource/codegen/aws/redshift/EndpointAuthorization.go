package redshift

type EndpointAuthorization struct {
	// Indicates whether to force the revoke action. If true, the Redshift-managed VPC endpoints associated with the endpoint authorization are also deleted. Default value is `false`.
	ForceDelete bool `json:"forceDelete,omitempty" yaml:"forceDelete,omitempty"`

	// The virtual private cloud (VPC) identifiers to grant access to. If none are specified all VPCs in shared account are allowed.
	VpcIds []string `json:"vpcIds,omitempty" yaml:"vpcIds,omitempty"`

	// The Amazon Web Services account ID to grant access to.
	Account string `json:"account,omitempty" yaml:"account,omitempty"`

	// The cluster identifier of the cluster to grant access to.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`
}
