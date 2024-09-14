package types

type Keyspaces_TablePointInTimeRecovery struct {
	// Valid values: `ENABLED`, `DISABLED`. The default value is `DISABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
