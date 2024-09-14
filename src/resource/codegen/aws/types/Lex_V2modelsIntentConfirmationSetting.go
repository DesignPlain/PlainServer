package types

type Lex_V2modelsIntentConfirmationSetting struct {
	// Configuration block for when the user answers "no" to the question defined in `prompt_specification`, Amazon Lex responds with this response to acknowledge that the intent was canceled. See `declination_response`.
	DeclinationResponse Lex_V2modelsIntentConfirmationSettingDeclinationResponse `json:"declinationResponse,omitempty" yaml:"declinationResponse,omitempty"`

	// Configuration block for when the code hook is invoked during confirmation prompt retries. See `elicitation_code_hook`.
	ElicitationCodeHook Lex_V2modelsIntentConfirmationSettingElicitationCodeHook `json:"elicitationCodeHook,omitempty" yaml:"elicitationCodeHook,omitempty"`

	// Configuration block for message groups that Amazon Lex uses to respond the user input. See `failure_response`.
	FailureResponse Lex_V2modelsIntentConfirmationSettingFailureResponse `json:"failureResponse,omitempty" yaml:"failureResponse,omitempty"`

	// Configuration block for prompting the user to confirm the intent. This question should have a yes or no answer. Amazon Lex uses this prompt to ensure that the user acknowledges that the intent is ready for fulfillment. See `prompt_specification`.
	PromptSpecification Lex_V2modelsIntentConfirmationSettingPromptSpecification `json:"promptSpecification,omitempty" yaml:"promptSpecification,omitempty"`

	// Configuration block for message groups that Amazon Lex uses to respond the user input. See `confirmation_response`.
	ConfirmationResponse Lex_V2modelsIntentConfirmationSettingConfirmationResponse `json:"confirmationResponse,omitempty" yaml:"confirmationResponse,omitempty"`

	// Configuration block for conditional branches to evaluate after the intent is declined. See `declination_conditional`.
	DeclinationConditional Lex_V2modelsIntentConfirmationSettingDeclinationConditional `json:"declinationConditional,omitempty" yaml:"declinationConditional,omitempty"`

	// Configuration block for conditional branches to evaluate after the intent is closed. See `confirmation_conditional`.
	ConfirmationConditional Lex_V2modelsIntentConfirmationSettingConfirmationConditional `json:"confirmationConditional,omitempty" yaml:"confirmationConditional,omitempty"`

	// Configuration block for the next step that the bot executes when the customer confirms the intent. See `confirmation_next_step`.
	ConfirmationNextStep Lex_V2modelsIntentConfirmationSettingConfirmationNextStep `json:"confirmationNextStep,omitempty" yaml:"confirmationNextStep,omitempty"`

	// Configuration block for the next step that the bot executes when the customer declines the intent. See `declination_next_step`.
	DeclinationNextStep Lex_V2modelsIntentConfirmationSettingDeclinationNextStep `json:"declinationNextStep,omitempty" yaml:"declinationNextStep,omitempty"`

	// Configuration block for conditional branches. Branches are evaluated in the order that they are entered in the list. The first branch with a condition that evaluates to true is executed. The last branch in the list is the default branch. The default branch should not have any condition expression. The default branch is executed if no other branch has a matching condition. See `failure_conditional`.
	FailureConditional Lex_V2modelsIntentConfirmationSettingFailureConditional `json:"failureConditional,omitempty" yaml:"failureConditional,omitempty"`

	// Configuration block for the next step to take in the conversation if the confirmation step fails. See `failure_next_step`.
	FailureNextStep Lex_V2modelsIntentConfirmationSettingFailureNextStep `json:"failureNextStep,omitempty" yaml:"failureNextStep,omitempty"`

	// Whether the intent's confirmation is sent to the user. When this field is false, confirmation and declination responses aren't sent. If the active field isn't specified, the default is true.
	Active bool `json:"active,omitempty" yaml:"active,omitempty"`

	// Configuration block for the intent's confirmation step. The dialog code hook is triggered based on these invocation settings when the confirmation next step or declination next step or failure next step is `invoke_dialog_code_hook`.  See `code_hook`.
	CodeHook Lex_V2modelsIntentConfirmationSettingCodeHook `json:"codeHook,omitempty" yaml:"codeHook,omitempty"`
}
