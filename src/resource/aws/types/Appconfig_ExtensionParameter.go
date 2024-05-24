package types

type Appconfig_ExtensionParameter struct {
	// Information about the parameter.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The parameter name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Determines if a parameter value must be specified in the extension association.
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`
}
