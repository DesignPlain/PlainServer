package types

type Backup_SelectionConditionStringEqual struct {
	// The value in a key-value pair.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// The key in a key-value pair.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
