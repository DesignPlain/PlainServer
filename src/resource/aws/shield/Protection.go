package shield

type Protection struct {
	// A friendly name for the Protection you are creating.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The ARN (Amazon Resource Name) of the resource to be protected.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
