package types

type Diagflow_CxPageFormParameter struct {
	/*
	   The entity type of the parameter.
	   Format: projects/-/locations/-/agents/-/entityTypes/<System Entity Type ID> for system entity types (for example, projects/-/locations/-/agents/-/entityTypes/sys.date), or projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/entityTypes/<Entity Type ID> for developer entity
	*/
	EntityType string `json:"entityType,omitempty" yaml:"entityType,omitempty"`

	/*
	   Defines fill behavior for the parameter.
	   Structure is documented below.
	*/
	FillBehavior Diagflow_CxPageFormParameterFillBehavior `json:"fillBehavior,omitempty" yaml:"fillBehavior,omitempty"`

	// Indicates whether the parameter represents a list of values.
	IsList bool `json:"isList,omitempty" yaml:"isList,omitempty"`

	/*
	   Indicates whether the parameter content should be redacted in log.
	   If redaction is enabled, the parameter content will be replaced by parameter name during logging. Note: the parameter content is subject to redaction if either parameter level redaction or entity type level redaction is enabled.
	*/
	Redact bool `json:"redact,omitempty" yaml:"redact,omitempty"`

	/*
	   Indicates whether the parameter is required. Optional parameters will not trigger prompts; however, they are filled if the user specifies them.
	   Required parameters must be filled before form filling concludes.
	*/
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`

	/*
	   Hierarchical advanced settings for this parameter. The settings exposed at the lower level overrides the settings exposed at the higher level.
	   Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	   Structure is documented below.
	*/
	AdvancedSettings Diagflow_CxPageFormParameterAdvancedSettings `json:"advancedSettings,omitempty" yaml:"advancedSettings,omitempty"`

	// The default value of an optional parameter. If the parameter is required, the default value will be ignored.
	DefaultValue string `json:"defaultValue,omitempty" yaml:"defaultValue,omitempty"`

	// The human-readable name of the parameter, unique within the form.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
