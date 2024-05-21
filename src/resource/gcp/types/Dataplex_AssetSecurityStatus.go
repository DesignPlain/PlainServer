package types

type Dataplex_AssetSecurityStatus struct {
	// Additional information about the current state.
	Message string `json:"message,omitempty" yaml:"message,omitempty"`

	// Output only. Current state of the asset. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// Output only. The time when the asset was last updated.
	UpdateTime string `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`
}
