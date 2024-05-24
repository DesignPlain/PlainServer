package types

type Lex_V2modelsIntentFulfillmentCodeHookFulfillmentUpdatesSpecification struct {
	// Whether fulfillment updates are sent to the user. When this field is true, updates are sent. If the active field is set to true, the `start_response`, `update_response`, and `timeout_in_seconds` fields are required.
	Active bool `json:"active,omitempty" yaml:"active,omitempty"`

	// Configuration block for the message sent to users when the fulfillment Lambda functions starts running.
	StartResponse Lex_V2modelsIntentFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponse `json:"startResponse,omitempty" yaml:"startResponse,omitempty"`

	// Length of time that the fulfillment Lambda function should run before it times out.
	TimeoutInSeconds int `json:"timeoutInSeconds,omitempty" yaml:"timeoutInSeconds,omitempty"`

	// Configuration block for messages sent periodically to the user while the fulfillment Lambda function is running.
	UpdateResponse Lex_V2modelsIntentFulfillmentCodeHookFulfillmentUpdatesSpecificationUpdateResponse `json:"updateResponse,omitempty" yaml:"updateResponse,omitempty"`
}
