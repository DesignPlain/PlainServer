package types

type Dataplex_LakeMetastoreStatus struct {
	// The URI of the endpoint used to access the Metastore service.
	Endpoint string `json:"endpoint,omitempty" yaml:"endpoint,omitempty"`

	// Additional information about the current status.
	Message string `json:"message,omitempty" yaml:"message,omitempty"`

	// Output only. Current state of the lake. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	// Output only. The time when the lake was last updated.
	UpdateTime string `json:"updateTime,omitempty" yaml:"updateTime,omitempty"`
}
