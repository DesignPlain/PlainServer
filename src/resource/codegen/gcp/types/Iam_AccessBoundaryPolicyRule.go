package types

type Iam_AccessBoundaryPolicyRule struct {
	/*
	   An access boundary rule in an IAM policy.
	   Structure is documented below.
	*/
	AccessBoundaryRule Iam_AccessBoundaryPolicyRuleAccessBoundaryRule `json:"accessBoundaryRule,omitempty" yaml:"accessBoundaryRule,omitempty"`

	// The description of the rule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
