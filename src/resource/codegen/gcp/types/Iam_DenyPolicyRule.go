package types

type Iam_DenyPolicyRule struct {
	/*
	   A deny rule in an IAM deny policy.
	   Structure is documented below.
	*/
	DenyRule Iam_DenyPolicyRuleDenyRule `json:"denyRule,omitempty" yaml:"denyRule,omitempty"`

	// The description of the rule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
