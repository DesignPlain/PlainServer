package cfg

type RecorderStatus struct {
	// Whether the configuration recorder should be enabled or disabled.
	IsEnabled bool `json:"isEnabled,omitempty" yaml:"isEnabled,omitempty"`

	// The name of the recorder
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
