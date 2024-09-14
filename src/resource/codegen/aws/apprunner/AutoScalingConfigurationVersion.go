package apprunner

type AutoScalingConfigurationVersion struct {
	// Name of the auto scaling configuration.
	AutoScalingConfigurationName string `json:"autoScalingConfigurationName,omitempty" yaml:"autoScalingConfigurationName,omitempty"`

	// Maximal number of concurrent requests that you want an instance to process. When the number of concurrent requests goes over this limit, App Runner scales up your service.
	MaxConcurrency int `json:"maxConcurrency,omitempty" yaml:"maxConcurrency,omitempty"`

	// Maximal number of instances that App Runner provisions for your service.
	MaxSize int `json:"maxSize,omitempty" yaml:"maxSize,omitempty"`

	// Minimal number of instances that App Runner provisions for your service.
	MinSize int `json:"minSize,omitempty" yaml:"minSize,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
