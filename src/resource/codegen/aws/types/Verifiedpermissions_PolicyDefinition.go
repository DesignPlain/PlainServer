package types

type Verifiedpermissions_PolicyDefinition struct {
	// The static policy statement. See Static below.
	Static Verifiedpermissions_PolicyDefinitionStatic `json:"static,omitempty" yaml:"static,omitempty"`

	// The template linked policy. See Template Linked below.
	TemplateLinked Verifiedpermissions_PolicyDefinitionTemplateLinked `json:"templateLinked,omitempty" yaml:"templateLinked,omitempty"`
}
