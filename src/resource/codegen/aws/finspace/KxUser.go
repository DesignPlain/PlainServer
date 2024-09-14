package finspace

type KxUser struct {
	// Unique identifier for the KX environment.
	EnvironmentId string `json:"environmentId,omitempty" yaml:"environmentId,omitempty"`

	/*
	   IAM role ARN to be associated with the user.

	   The following arguments are optional:
	*/
	IamRole string `json:"iamRole,omitempty" yaml:"iamRole,omitempty"`

	// A unique identifier for the user.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
