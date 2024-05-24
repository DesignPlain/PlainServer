package types

type Budgets_BudgetActionDefinitionSsmActionDefinition struct {
	// The Region to run the SSM document.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The action subType. Valid values are `STOP_EC2_INSTANCES` or `STOP_RDS_INSTANCES`.
	ActionSubType string `json:"actionSubType,omitempty" yaml:"actionSubType,omitempty"`

	// The EC2 and RDS instance IDs.
	InstanceIds []string `json:"instanceIds,omitempty" yaml:"instanceIds,omitempty"`
}
