package types

type Lex_V2modelsIntentFulfillmentCodeHookPostFulfillmentStatusSpecification struct {
	// Configuration block for the next step the bot runs after the dialog code hook finishes successfully. See `success_next_step`.
	SuccessNextStep Lex_V2modelsIntentFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStep `json:"successNextStep,omitempty" yaml:"successNextStep,omitempty"`

	// Configuration block for the next step that the bot runs when the code hook times out. See `timeout_next_step`.
	TimeoutNextStep Lex_V2modelsIntentFulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutNextStep `json:"timeoutNextStep,omitempty" yaml:"timeoutNextStep,omitempty"`

	// Configuration block for message groups that Amazon Lex uses to respond the user input. See `failure_response`.
	FailureResponse Lex_V2modelsIntentFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponse `json:"failureResponse,omitempty" yaml:"failureResponse,omitempty"`

	// Configuration block for conditional branches to evaluate after the dialog code hook finishes successfully. See `success_conditional`.
	SuccessConditional Lex_V2modelsIntentFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditional `json:"successConditional,omitempty" yaml:"successConditional,omitempty"`

	// Configuration block for message groups that Amazon Lex uses to respond the user input. See `success_response`.
	SuccessResponse Lex_V2modelsIntentFulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponse `json:"successResponse,omitempty" yaml:"successResponse,omitempty"`

	// Configuration block for conditional branches to evaluate if the code hook times out. See `timeout_conditional`.
	TimeoutConditional Lex_V2modelsIntentFulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditional `json:"timeoutConditional,omitempty" yaml:"timeoutConditional,omitempty"`

	// Configuration block for a list of message groups that Amazon Lex uses to respond the user input. See `timeout_response`.
	TimeoutResponse Lex_V2modelsIntentFulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutResponse `json:"timeoutResponse,omitempty" yaml:"timeoutResponse,omitempty"`

	// Configuration block for conditional branches to evaluate after the dialog code hook throws an exception or returns with the State field of the Intent object set to Failed. See `failure_conditional`.
	FailureConditional Lex_V2modelsIntentFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureConditional `json:"failureConditional,omitempty" yaml:"failureConditional,omitempty"`

	// Configuration block for the next step the bot runs after the dialog code hook throws an exception or returns with the State field of the Intent object set to Failed. See `failure_next_step`.
	FailureNextStep Lex_V2modelsIntentFulfillmentCodeHookPostFulfillmentStatusSpecificationFailureNextStep `json:"failureNextStep,omitempty" yaml:"failureNextStep,omitempty"`
}
