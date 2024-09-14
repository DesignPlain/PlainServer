package emr

type Studio struct {
	// The IAM user role that users and groups assume when logged in to an Amazon EMR Studio. Only specify a User Role when you use Amazon Web Services SSO authentication. The permissions attached to the User Role can be scoped down for each user or group using session policies.
	UserRole string `json:"userRole,omitempty" yaml:"userRole,omitempty"`

	// The ID of the Amazon Virtual Private Cloud (Amazon VPC) to associate with the Studio.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	/*
	   The ID of the Amazon EMR Studio Workspace security group. The Workspace security group allows outbound network traffic to resources in the Engine security group, and it must be in the same VPC specified by `vpc_id`.

	   The following arguments are optional:
	*/
	WorkspaceSecurityGroupId string `json:"workspaceSecurityGroupId,omitempty" yaml:"workspaceSecurityGroupId,omitempty"`

	// A detailed description of the Amazon EMR Studio.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The authentication endpoint of your identity provider (IdP). Specify this value when you use IAM authentication and want to let federated users log in to a Studio with the Studio URL and credentials from your IdP. Amazon EMR Studio redirects users to this endpoint to enter credentials.
	IdpAuthUrl string `json:"idpAuthUrl,omitempty" yaml:"idpAuthUrl,omitempty"`

	// The name that your identity provider (IdP) uses for its RelayState parameter. For example, RelayState or TargetSource. Specify this value when you use IAM authentication and want to let federated users log in to a Studio using the Studio URL. The RelayState parameter differs by IdP.
	IdpRelayStateParameterName string `json:"idpRelayStateParameterName,omitempty" yaml:"idpRelayStateParameterName,omitempty"`

	// The IAM role that the Amazon EMR Studio assumes. The service role provides a way for Amazon EMR Studio to interoperate with other Amazon Web Services services.
	ServiceRole string `json:"serviceRole,omitempty" yaml:"serviceRole,omitempty"`

	// A list of subnet IDs to associate with the Amazon EMR Studio. A Studio can have a maximum of 5 subnets. The subnets must belong to the VPC specified by `vpc_id`. Studio users can create a Workspace in any of the specified subnets.
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// list of tags to apply to the EMR Cluster. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies whether the Studio authenticates users using IAM or Amazon Web Services SSO. Valid values are `SSO` or `IAM`.
	AuthMode string `json:"authMode,omitempty" yaml:"authMode,omitempty"`

	// The Amazon S3 location to back up Amazon EMR Studio Workspaces and notebook files.
	DefaultS3Location string `json:"defaultS3Location,omitempty" yaml:"defaultS3Location,omitempty"`

	// The ID of the Amazon EMR Studio Engine security group. The Engine security group allows inbound network traffic from the Workspace security group, and it must be in the same VPC specified by `vpc_id`.
	EngineSecurityGroupId string `json:"engineSecurityGroupId,omitempty" yaml:"engineSecurityGroupId,omitempty"`

	// A descriptive name for the Amazon EMR Studio.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
