package types

type Compute_getInstanceGroupManagerVersion struct {
	// The full URL to an instance template from which all new instances of this version will be created.
	InstanceTemplate string `json:"instanceTemplate,omitempty" yaml:"instanceTemplate,omitempty"`

	// The name of the instance group. Either `name` or `self_link` must be provided.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The number of instances calculated as a fixed number or a percentage depending on the settings.
	TargetSizes []Compute_getInstanceGroupManagerVersionTargetSize `json:"targetSizes,omitempty" yaml:"targetSizes,omitempty"`
}
