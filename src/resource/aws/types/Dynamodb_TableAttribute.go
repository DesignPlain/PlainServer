package types

type Dynamodb_TableAttribute struct {
	// Name of the attribute
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Attribute type. Valid values are `S` (string), `N` (number), `B` (binary).
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
