package types

type Dataproc_WorkflowTemplateParameterValidationRegex struct {
	// Required. RE2 regular expressions used to validate the parameter's value. The value must match the regex in its entirety (substring matches are not sufficient).
	Regexes []string `json:"regexes,omitempty" yaml:"regexes,omitempty"`
}
