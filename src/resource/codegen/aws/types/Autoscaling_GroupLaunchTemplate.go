package types

type Autoscaling_GroupLaunchTemplate struct {
	// ID of the launch template. Conflicts with `name`.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Name of the launch template. Conflicts with `id`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Template version. Can be version number, `$Latest`, or `$Default`. (Default: `$Default`).
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
