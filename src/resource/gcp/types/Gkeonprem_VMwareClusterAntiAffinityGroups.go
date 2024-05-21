package types

type Gkeonprem_VMwareClusterAntiAffinityGroups struct {
	/*
	   Spread nodes across at least three physical hosts (requires at least three
	   hosts).
	   Enabled by default.
	*/
	AagConfigDisabled bool `json:"aagConfigDisabled,omitempty" yaml:"aagConfigDisabled,omitempty"`
}
