package types

type Verifiedpermissions_PolicyDefinitionTemplateLinked struct {
	// The ID of the template.
	PolicyTemplateId string `json:"policyTemplateId,omitempty" yaml:"policyTemplateId,omitempty"`

	// The principal of the template linked policy.
	Principal Verifiedpermissions_PolicyDefinitionTemplateLinkedPrincipal `json:"principal,omitempty" yaml:"principal,omitempty"`

	// The resource of the template linked policy.
	Resource Verifiedpermissions_PolicyDefinitionTemplateLinkedResource `json:"resource,omitempty" yaml:"resource,omitempty"`
}
