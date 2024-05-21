package types

type Dataloss_PreventionJobTriggerInspectJobActionDeidentifyTransformationConfig struct {
	// If this template is specified, it will serve as the default de-identify template.
	DeidentifyTemplate string `json:"deidentifyTemplate,omitempty" yaml:"deidentifyTemplate,omitempty"`

	// If this template is specified, it will serve as the de-identify template for images.
	ImageRedactTemplate string `json:"imageRedactTemplate,omitempty" yaml:"imageRedactTemplate,omitempty"`

	// If this template is specified, it will serve as the de-identify template for structured content such as delimited files and tables.
	StructuredDeidentifyTemplate string `json:"structuredDeidentifyTemplate,omitempty" yaml:"structuredDeidentifyTemplate,omitempty"`
}
