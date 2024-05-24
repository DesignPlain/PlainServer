package types

type Autoscaling_getGroupLaunchTemplate struct {
	// ID of the launch template.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Specify the exact name of the desired autoscaling group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Template version.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
