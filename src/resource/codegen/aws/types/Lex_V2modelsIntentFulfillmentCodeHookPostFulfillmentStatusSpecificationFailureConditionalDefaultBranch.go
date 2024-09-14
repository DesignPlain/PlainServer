package types

type Lex_V2modelsIntentFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureConditionalDefaultBranch struct {
	// Configuration block for the next step in the conversation. See `next_step`.
	NextStep Lex_V2modelsIntentFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureConditionalDefaultBranchNextStep `json:"nextStep,omitempty" yaml:"nextStep,omitempty"`

	// Configuration block for a list of message groups that Amazon Lex uses to respond to the user input. See `response`.
	Response Lex_V2modelsIntentFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureConditionalDefaultBranchResponse `json:"response,omitempty" yaml:"response,omitempty"`
}
