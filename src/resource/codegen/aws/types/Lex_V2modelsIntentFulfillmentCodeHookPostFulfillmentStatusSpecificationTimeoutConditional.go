package types

type Lex_V2modelsIntentFulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditional struct {
	// Configuration blocks for conditional branches. A conditional branch is made up of a condition, a response and a next step. The response and next step are executed when the condition is true. See `conditional_branch`.
	ConditionalBranches []Lex_V2modelsIntentFulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranch `json:"conditionalBranches,omitempty" yaml:"conditionalBranches,omitempty"`

	// Configuration block for the conditional branch that should be followed when the conditions for other branches are not satisfied. A branch is made up of a condition, a response and a next step. See `default_branch`.
	DefaultBranch Lex_V2modelsIntentFulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalDefaultBranch `json:"defaultBranch,omitempty" yaml:"defaultBranch,omitempty"`

	// Whether a conditional branch is active. When active is false, the conditions are not evaluated.
	Active bool `json:"active,omitempty" yaml:"active,omitempty"`
}
