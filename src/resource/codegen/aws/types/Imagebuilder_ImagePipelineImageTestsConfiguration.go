package types

type Imagebuilder_ImagePipelineImageTestsConfiguration struct {
	// Whether image tests are enabled. Defaults to `true`.
	ImageTestsEnabled bool `json:"imageTestsEnabled,omitempty" yaml:"imageTestsEnabled,omitempty"`

	// Number of minutes before image tests time out. Valid values are between `60` and `1440`. Defaults to `720`.
	TimeoutMinutes int `json:"timeoutMinutes,omitempty" yaml:"timeoutMinutes,omitempty"`
}
