package types

type Securitycenter_SourceIamBindingCondition struct {
	// The description of the source (max of 1024 characters).
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	//
	Expression string `json:"expression,omitempty" yaml:"expression,omitempty"`

	//
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
}
