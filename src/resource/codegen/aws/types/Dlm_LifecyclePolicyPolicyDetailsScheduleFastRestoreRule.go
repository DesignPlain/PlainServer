package types

type Dlm_LifecyclePolicyPolicyDetailsScheduleFastRestoreRule struct {
	// The Availability Zones in which to enable fast snapshot restore.
	AvailabilityZones []string `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`

	//
	Count int `json:"count,omitempty" yaml:"count,omitempty"`

	//
	Interval int `json:"interval,omitempty" yaml:"interval,omitempty"`

	//
	IntervalUnit string `json:"intervalUnit,omitempty" yaml:"intervalUnit,omitempty"`
}
