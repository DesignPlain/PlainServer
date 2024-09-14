package serverlessrepository

type CloudFormationStack struct {
	// The name of the stack to create. The resource deployed in AWS will be prefixed with `serverlessrepo-`
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of Parameter structures that specify input parameters for the stack.
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// The version of the application to deploy. If not supplied, deploys the latest version.
	SemanticVersion string `json:"semanticVersion,omitempty" yaml:"semanticVersion,omitempty"`

	// A list of tags to associate with this stack. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ARN of the application from the Serverless Application Repository.
	ApplicationId string `json:"applicationId,omitempty" yaml:"applicationId,omitempty"`

	// A list of capabilities. Valid values are `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_RESOURCE_POLICY`, or `CAPABILITY_AUTO_EXPAND`
	Capabilities []string `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`
}
