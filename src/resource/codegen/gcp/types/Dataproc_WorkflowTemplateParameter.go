package types

type Dataproc_WorkflowTemplateParameter struct {
	// Validation rules to be applied to this parameter's value.
	Validation Dataproc_WorkflowTemplateParameterValidation `json:"validation,omitempty" yaml:"validation,omitempty"`

	// Brief description of the parameter. Must not exceed 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Required. Paths to all fields that the parameter replaces. A field is allowed to appear in at most one parameter's list of field paths. A field path is similar in syntax to a .sparkJob.args
	Fields []string `json:"fields,omitempty" yaml:"fields,omitempty"`

	// Required. Parameter name. The parameter name is used as the key, and paired with the parameter value, which are passed to the template when the template is instantiated. The name must contain only capital letters (A-Z), numbers (0-9), and underscores (_), and must not start with a number. The maximum length is 40 characters.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
