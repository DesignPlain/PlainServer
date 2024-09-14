package verifiedaccess

type Instance struct {
	// A description for the AWS Verified Access Instance.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Enable or disable support for Federal Information Processing Standards (FIPS) on the AWS Verified Access Instance.
	FipsEnabled bool `json:"fipsEnabled,omitempty" yaml:"fipsEnabled,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
