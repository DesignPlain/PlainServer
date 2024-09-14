package schemas

type Discoverer struct {
	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The description of the discoverer. Maximum of 256 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The ARN of the event bus to discover event schemas on.
	SourceArn string `json:"sourceArn,omitempty" yaml:"sourceArn,omitempty"`
}
