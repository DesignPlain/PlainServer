package types

type Integrationconnectors_ConnectionLockConfig struct {
	// Indicates whether or not the connection is locked.
	Locked bool `json:"locked,omitempty" yaml:"locked,omitempty"`

	// Describes why a connection is locked.
	Reason string `json:"reason,omitempty" yaml:"reason,omitempty"`
}
