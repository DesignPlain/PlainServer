package codecommit

type ApprovalRuleTemplate struct {
	// The content of the approval rule template. Maximum of 3000 characters.
	Content string `json:"content,omitempty" yaml:"content,omitempty"`

	// The description of the approval rule template. Maximum of 1000 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name for the approval rule template. Maximum of 100 characters.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
