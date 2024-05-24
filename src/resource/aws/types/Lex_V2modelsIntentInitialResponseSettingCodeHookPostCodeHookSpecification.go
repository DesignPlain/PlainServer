package types

type Lex_V2modelsIntentInitialResponseSettingCodeHookPostCodeHookSpecification struct {
	// Configuration block for conditional branches to evaluate after the dialog code hook throws an exception or returns with the State field of the Intent object set to Failed.
	FailureConditional Lex_V2modelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureConditional `json:"failureConditional,omitempty" yaml:"failureConditional,omitempty"`

	// Configuration block for message groups that Amazon Lex uses to respond the user input. See `failure_response`.
	FailureResponse Lex_V2modelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureResponse `json:"failureResponse,omitempty" yaml:"failureResponse,omitempty"`

	// Configuration block for message groups that Amazon Lex uses to respond the user input. See `success_response`.
	SuccessResponse Lex_V2modelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationSuccessResponse `json:"successResponse,omitempty" yaml:"successResponse,omitempty"`

	// Configuration block for the next step the bot runs after the dialog code hook throws an exception or returns with the State field of the Intent object set to Failed . See `failure_next_step`.
	FailureNextStep Lex_V2modelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureNextStep `json:"failureNextStep,omitempty" yaml:"failureNextStep,omitempty"`

	// Configuration block for conditional branches to evaluate after the dialog code hook finishes successfully. See `success_conditional`.
	SuccessConditional Lex_V2modelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationSuccessConditional `json:"successConditional,omitempty" yaml:"successConditional,omitempty"`

	// Configuration block for the next step the bot runs after the dialog code hook finishes successfully. See `success_next_step`.
	SuccessNextStep Lex_V2modelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStep `json:"successNextStep,omitempty" yaml:"successNextStep,omitempty"`

	// Configuration block for conditional branches to evaluate if the code hook times out. See `timeout_conditional`.
	TimeoutConditional Lex_V2modelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationTimeoutConditional `json:"timeoutConditional,omitempty" yaml:"timeoutConditional,omitempty"`

	// Configuration block for the next step that the bot runs when the code hook times out. See `timeout_next_step`.
	TimeoutNextStep Lex_V2modelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationTimeoutNextStep `json:"timeoutNextStep,omitempty" yaml:"timeoutNextStep,omitempty"`

	// Configuration block for a list of message groups that Amazon Lex uses to respond the user input. See `timeout_response`.
	TimeoutResponse Lex_V2modelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationTimeoutResponse `json:"timeoutResponse,omitempty" yaml:"timeoutResponse,omitempty"`
}
