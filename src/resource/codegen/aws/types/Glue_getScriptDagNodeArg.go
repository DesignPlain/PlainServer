package types

type Glue_getScriptDagNodeArg struct {
	// Name of the argument or property.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Boolean if the value is used as a parameter. Defaults to `false`.
	Param bool `json:"param,omitempty" yaml:"param,omitempty"`

	// Value of the argument or property.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
