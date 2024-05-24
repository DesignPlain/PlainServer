package vpclattice

type ServiceNetworkVpcAssociation struct {
	// The IDs of the security groups.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	/*
	   The ID or Amazon Resource Identifier (ARN) of the service network. You must use the ARN if the resources specified in the operation are in different accounts.
	   The following arguments are optional:
	*/
	ServiceNetworkIdentifier string `json:"serviceNetworkIdentifier,omitempty" yaml:"serviceNetworkIdentifier,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ID of the VPC.
	VpcIdentifier string `json:"vpcIdentifier,omitempty" yaml:"vpcIdentifier,omitempty"`
}
