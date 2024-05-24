package types

type Lex_V2modelsIntentFulfillmentCodeHook struct {
	// Whether the fulfillment code hook is used. When active is false, the code hook doesn't run.
	Active bool `json:"active,omitempty" yaml:"active,omitempty"`

	// Whether a Lambda function should be invoked to fulfill a specific intent.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Configuration block for settings for update messages sent to the user for long-running Lambda fulfillment functions. Fulfillment updates can be used only with streaming conversations. See `fulfillment_updates_specification`.
	FulfillmentUpdatesSpecification Lex_V2modelsIntentFulfillmentCodeHookFulfillmentUpdatesSpecification `json:"fulfillmentUpdatesSpecification,omitempty" yaml:"fulfillmentUpdatesSpecification,omitempty"`

	// Configuration block for settings for messages sent to the user for after the Lambda fulfillment function completes. Post-fulfillment messages can be sent for both streaming and non-streaming conversations. See `post_fulfillment_status_specification`.
	PostFulfillmentStatusSpecification Lex_V2modelsIntentFulfillmentCodeHookPostFulfillmentStatusSpecification `json:"postFulfillmentStatusSpecification,omitempty" yaml:"postFulfillmentStatusSpecification,omitempty"`
}
