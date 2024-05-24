package types

type Imagebuilder_getImagePipelineImageTestsConfiguration struct {
	// Whether image tests are enabled.
	ImageTestsEnabled bool `json:"imageTestsEnabled,omitempty" yaml:"imageTestsEnabled,omitempty"`

	// Number of minutes before image tests time out.
	TimeoutMinutes int `json:"timeoutMinutes,omitempty" yaml:"timeoutMinutes,omitempty"`
}
