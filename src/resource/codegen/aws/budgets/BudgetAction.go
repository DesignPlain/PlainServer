package budgets

import types "libds/aws/types"

type BudgetAction struct {
	// The type of action. This defines the type of tasks that can be carried out by this action. This field also determines the format for definition. Valid values are `APPLY_IAM_POLICY`, `APPLY_SCP_POLICY`, and `RUN_SSM_DOCUMENTS`.
	ActionType string `json:"actionType,omitempty" yaml:"actionType,omitempty"`

	// The name of a budget.
	BudgetName string `json:"budgetName,omitempty" yaml:"budgetName,omitempty"`

	// Specifies all of the type-specific parameters. See Definition.
	Definition types.Budgets_BudgetActionDefinition `json:"definition,omitempty" yaml:"definition,omitempty"`

	// Map of tags assigned to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The type of a notification. Valid values are `ACTUAL` or `FORECASTED`.
	NotificationType string `json:"notificationType,omitempty" yaml:"notificationType,omitempty"`

	// A list of subscribers. See Subscriber.
	Subscribers []types.Budgets_BudgetActionSubscriber `json:"subscribers,omitempty" yaml:"subscribers,omitempty"`

	// The ID of the target account for budget. Will use current user's account_id by default if omitted.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// The trigger threshold of the action. See Action Threshold.
	ActionThreshold types.Budgets_BudgetActionActionThreshold `json:"actionThreshold,omitempty" yaml:"actionThreshold,omitempty"`

	// This specifies if the action needs manual or automatic approval. Valid values are `AUTOMATIC` and `MANUAL`.
	ApprovalModel string `json:"approvalModel,omitempty" yaml:"approvalModel,omitempty"`

	// The role passed for action execution and reversion. Roles and actions must be in the same account.
	ExecutionRoleArn string `json:"executionRoleArn,omitempty" yaml:"executionRoleArn,omitempty"`
}
