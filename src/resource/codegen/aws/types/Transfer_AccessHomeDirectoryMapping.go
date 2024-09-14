package types

type Transfer_AccessHomeDirectoryMapping struct {
	// Represents an entry and a target.
	Entry string `json:"entry,omitempty" yaml:"entry,omitempty"`

	// Represents the map target.
	Target string `json:"target,omitempty" yaml:"target,omitempty"`
}
