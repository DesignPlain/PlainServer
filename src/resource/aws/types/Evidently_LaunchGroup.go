package types

type Evidently_LaunchGroup struct {
	// Specifies the description of the launch group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Specifies the name of the feature that the launch is using.
	Feature string `json:"feature,omitempty" yaml:"feature,omitempty"`

	// Specifies the name of the lahnch group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies the feature variation to use for this launch group.
	Variation string `json:"variation,omitempty" yaml:"variation,omitempty"`
}
