package types

type Lex_IntentFulfillmentActivityCodeHook struct {
	/*
	   The version of the request-response that you want Amazon Lex to use
	   to invoke your Lambda function. For more information, see
	   [Using Lambda Functions](https://docs.aws.amazon.com/lex/latest/dg/using-lambda.html). Must be less than or equal to 5 characters in length.
	*/
	MessageVersion string `json:"messageVersion,omitempty" yaml:"messageVersion,omitempty"`

	// The Amazon Resource Name (ARN) of the Lambda function.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`
}
