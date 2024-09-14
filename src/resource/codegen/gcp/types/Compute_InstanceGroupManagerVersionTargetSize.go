package types

type Compute_InstanceGroupManagerVersionTargetSize struct {
	// , The number of instances which are managed for this version. Conflicts with `percent`.
	Fixed int `json:"fixed,omitempty" yaml:"fixed,omitempty"`

	/*
	   , The number of instances (calculated as percentage) which are managed for this version. Conflicts with `fixed`.
	   Note that when using `percent`, rounding will be in favor of explicitly set `target_size` values; a managed instance group with 2 instances and 2 `version`s,
	   one of which has a `target_size.percent` of `60` will create 2 instances of that `version`.
	*/
	Percent int `json:"percent,omitempty" yaml:"percent,omitempty"`
}
