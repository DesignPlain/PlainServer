package types

type Ssm_DocumentParameter struct {
	//
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	//
	DefaultValue string `json:"defaultValue,omitempty" yaml:"defaultValue,omitempty"`

	// The description of the document.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the document.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
