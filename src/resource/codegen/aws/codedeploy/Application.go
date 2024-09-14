package codedeploy

type Application struct {
	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The compute platform can either be `ECS`, `Lambda`, or `Server`. Default is `Server`.
	ComputePlatform string `json:"computePlatform,omitempty" yaml:"computePlatform,omitempty"`

	// The name of the application.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
