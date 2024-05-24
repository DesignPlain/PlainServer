package types

type Budgets_BudgetActionDefinitionScpActionDefinition struct {
	// The policy ID attached.
	PolicyId string `json:"policyId,omitempty" yaml:"policyId,omitempty"`

	// A list of target IDs.
	TargetIds []string `json:"targetIds,omitempty" yaml:"targetIds,omitempty"`
}
