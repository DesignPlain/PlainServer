package types

type Verifiedpermissions_PolicyDefinitionStatic struct {
	// The description of the static policy.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The statement of the static policy.
	Statement string `json:"statement,omitempty" yaml:"statement,omitempty"`
}
