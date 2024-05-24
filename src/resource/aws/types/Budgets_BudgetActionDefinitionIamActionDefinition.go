package types

type Budgets_BudgetActionDefinitionIamActionDefinition struct {
	// A list of groups to be attached. There must be at least one group.
	Groups []string `json:"groups,omitempty" yaml:"groups,omitempty"`

	// The Amazon Resource Name (ARN) of the policy to be attached.
	PolicyArn string `json:"policyArn,omitempty" yaml:"policyArn,omitempty"`

	// A list of roles to be attached. There must be at least one role.
	Roles []string `json:"roles,omitempty" yaml:"roles,omitempty"`

	// A list of users to be attached. There must be at least one user.
	Users []string `json:"users,omitempty" yaml:"users,omitempty"`
}
