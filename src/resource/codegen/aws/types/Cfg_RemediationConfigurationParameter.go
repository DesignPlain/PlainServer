package types

type Cfg_RemediationConfigurationParameter struct {
	// Name of the attribute.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Value is dynamic and changes at run-time.
	ResourceValue string `json:"resourceValue,omitempty" yaml:"resourceValue,omitempty"`

	// Value is static and does not change at run-time.
	StaticValue string `json:"staticValue,omitempty" yaml:"staticValue,omitempty"`

	// List of static values.
	StaticValues []string `json:"staticValues,omitempty" yaml:"staticValues,omitempty"`
}
