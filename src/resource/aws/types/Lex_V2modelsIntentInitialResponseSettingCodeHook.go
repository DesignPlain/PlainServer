package types

type Lex_V2modelsIntentInitialResponseSettingCodeHook struct {
	// Label that indicates the dialog step from which the dialog code hook is happening.
	InvocationLabel string `json:"invocationLabel,omitempty" yaml:"invocationLabel,omitempty"`

	// Configuration block that contains the responses and actions that Amazon Lex takes after the Lambda function is complete. See `post_code_hook_specification`.
	PostCodeHookSpecification Lex_V2modelsIntentInitialResponseSettingCodeHookPostCodeHookSpecification `json:"postCodeHookSpecification,omitempty" yaml:"postCodeHookSpecification,omitempty"`

	// Whether a dialog code hook is used when the intent is activated.
	Active bool `json:"active,omitempty" yaml:"active,omitempty"`

	// Whether a Lambda function should be invoked for the dialog.
	EnableCodeHookInvocation bool `json:"enableCodeHookInvocation,omitempty" yaml:"enableCodeHookInvocation,omitempty"`
}
