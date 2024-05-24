package types

type Keyspaces_TableClientSideTimestamps struct {
	// Shows how to enable client-side timestamps settings for the specified table. Valid values: `ENABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
