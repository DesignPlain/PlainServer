package types

type Lex_IntentFulfillmentActivity struct {
	/*
	   A description of the Lambda function that is run to fulfill the intent.
	   Required if type is CodeHook. Attributes are documented under code_hook.
	*/
	CodeHook Lex_IntentFulfillmentActivityCodeHook `json:"codeHook,omitempty" yaml:"codeHook,omitempty"`

	/*
	   How the intent should be fulfilled, either by running a Lambda function or by
	   returning the slot data to the client application. Type can be either `ReturnIntent` or `CodeHook`, as documented [here](https://docs.aws.amazon.com/lex/latest/dg/API_FulfillmentActivity.html).
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
