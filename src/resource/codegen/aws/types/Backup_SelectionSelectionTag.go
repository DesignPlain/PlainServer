package types

type Backup_SelectionSelectionTag struct {
	// The key in a key-value pair.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// An operation, such as `StringEquals`, that is applied to a key-value pair used to filter resources in a selection.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The value in a key-value pair.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
