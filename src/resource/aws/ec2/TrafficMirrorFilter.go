package ec2

type TrafficMirrorFilter struct {
	// A description of the filter.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// List of amazon network services that should be mirrored. Valid values: `amazon-dns`.
	NetworkServices []string `json:"networkServices,omitempty" yaml:"networkServices,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
