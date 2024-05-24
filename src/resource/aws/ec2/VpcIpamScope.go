package ec2

type VpcIpamScope struct {
	// A description for the scope you're creating.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The ID of the IPAM for which you're creating this scope.
	IpamId string `json:"ipamId,omitempty" yaml:"ipamId,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
