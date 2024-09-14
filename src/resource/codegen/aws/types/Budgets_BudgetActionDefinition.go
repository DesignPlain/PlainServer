package types

type Budgets_BudgetActionDefinition struct {
	// The service control policies (SCPs) action definition details. See SCP Action Definition.
	ScpActionDefinition Budgets_BudgetActionDefinitionScpActionDefinition `json:"scpActionDefinition,omitempty" yaml:"scpActionDefinition,omitempty"`

	// The AWS Systems Manager (SSM) action definition details. See SSM Action Definition.
	SsmActionDefinition Budgets_BudgetActionDefinitionSsmActionDefinition `json:"ssmActionDefinition,omitempty" yaml:"ssmActionDefinition,omitempty"`

	// The AWS Identity and Access Management (IAM) action definition details. See IAM Action Definition.
	IamActionDefinition Budgets_BudgetActionDefinitionIamActionDefinition `json:"iamActionDefinition,omitempty" yaml:"iamActionDefinition,omitempty"`
}
