package vpclattice

type ServiceNetwork struct {
	// Type of IAM policy. Either `NONE` or `AWS_IAM`.
	AuthType string `json:"authType,omitempty" yaml:"authType,omitempty"`

	/*
	   Name of the service network

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
