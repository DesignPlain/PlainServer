package types

type Imagebuilder_ImageWorkflowParameter struct {
	// The name of the Workflow parameter.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The value of the Workflow parameter.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
