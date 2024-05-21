package types

type Dataproc_WorkflowTemplateParameterValidation struct {
	// Validation based on a list of allowed values.
	Values Dataproc_WorkflowTemplateParameterValidationValues `json:"values,omitempty" yaml:"values,omitempty"`

	// Validation based on regular expressions.
	Regex Dataproc_WorkflowTemplateParameterValidationRegex `json:"regex,omitempty" yaml:"regex,omitempty"`
}
