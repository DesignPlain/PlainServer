package rds

type ProxyEndpoint struct {
	// One or more VPC security group IDs to associate with the new proxy.
	VpcSecurityGroupIds []string `json:"vpcSecurityGroupIds,omitempty" yaml:"vpcSecurityGroupIds,omitempty"`

	// One or more VPC subnet IDs to associate with the new proxy.
	VpcSubnetIds []string `json:"vpcSubnetIds,omitempty" yaml:"vpcSubnetIds,omitempty"`

	// The identifier for the proxy endpoint. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens; it can't end with a hyphen or contain two consecutive hyphens.
	DbProxyEndpointName string `json:"dbProxyEndpointName,omitempty" yaml:"dbProxyEndpointName,omitempty"`

	// The name of the DB proxy associated with the DB proxy endpoint that you create.
	DbProxyName string `json:"dbProxyName,omitempty" yaml:"dbProxyName,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Indicates whether the DB proxy endpoint can be used for read/write or read-only operations. The default is `READ_WRITE`. Valid values are `READ_WRITE` and `READ_ONLY`.
	TargetRole string `json:"targetRole,omitempty" yaml:"targetRole,omitempty"`
}
