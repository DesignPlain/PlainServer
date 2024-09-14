package types

type Ssm_DocumentParameter struct {
	// If specified, the default values for the parameters. Parameters without a default value are required. Parameters with a default value are optional.
	DefaultValue string `json:"defaultValue,omitempty" yaml:"defaultValue,omitempty"`

	// A description of what the parameter does, how to use it, the default value, and whether or not the parameter is optional.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the document.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The type of parameter. Valid values: `String`, `StringList`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
