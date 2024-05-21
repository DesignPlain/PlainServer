package types

type Compute_InstanceSettingsMetadata struct {
	// A metadata key/value items map. The total size of all keys and values must be less than 512KB
	Items map[string]string `json:"items,omitempty" yaml:"items,omitempty"`
}
