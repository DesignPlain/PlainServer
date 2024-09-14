package types

type Dataplex_AssetDiscoveryStatus struct {
	// The duration of the last discovery run.
	LastRunDuration string `json:"lastRunDuration,omitempty" yaml:"lastRunDuration,omitempty"`

	// The start time of the last discovery run.
	LastRunTime string `json:"lastRunTime,omitempty" yaml:"lastRunTime,omitempty"`

	// Additional information about the current state.
	Message string `json:"message,omitempty" yaml:"message,omitempty"`

	// Output only. Current state of the asset. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// Data Stats of the asset reported by discovery.
	Stats []Dataplex_AssetDiscoveryStatusStat `json:"stats,omitempty" yaml:"stats,omitempty"`

	// Output only. The time when the asset was last updated.
	UpdateTime string `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`
}
