package types

type Compute_InstanceGroupManagerVersion struct {
	// The full URL to an instance template from which all new instances of this version will be created. It is recommended to reference instance templates through their unique id (`self_link_unique` attribute).
	InstanceTemplate string `json:"instanceTemplate,omitempty" yaml:"instanceTemplate,omitempty"`

	// Version name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.

	   > Exactly one `version` you specify must not have a `target_size` specified. During a rolling update, the instance group manager will fulfill the `target_size`
	   constraints of every other `version`, and any remaining instances will be provisioned with the version where `target_size` is unset.
	*/
	TargetSize Compute_InstanceGroupManagerVersionTargetSize `json:"targetSize,omitempty" yaml:"targetSize,omitempty"`
}
