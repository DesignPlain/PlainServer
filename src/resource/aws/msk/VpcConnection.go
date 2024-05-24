package msk

type VpcConnection struct {
	// The VPC ID of the remote client.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// The authentication type for the client VPC connection. Specify one of these auth type strings: SASL_IAM, SASL_SCRAM, or TLS.
	Authentication string `json:"authentication,omitempty" yaml:"authentication,omitempty"`

	// The list of subnets in the client VPC to connect to.
	ClientSubnets []string `json:"clientSubnets,omitempty" yaml:"clientSubnets,omitempty"`

	// The security groups to attach to the ENIs for the broker nodes.
	SecurityGroups []string `json:"securityGroups,omitempty" yaml:"securityGroups,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The Amazon Resource Name (ARN) of the cluster.
	TargetClusterArn string `json:"targetClusterArn,omitempty" yaml:"targetClusterArn,omitempty"`
}
