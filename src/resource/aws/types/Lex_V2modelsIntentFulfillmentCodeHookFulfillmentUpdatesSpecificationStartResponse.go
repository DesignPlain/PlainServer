package types

type Lex_V2modelsIntentFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponse struct {
	// Whether the user can interrupt the start message while it is playing.
	AllowInterrupt bool `json:"allowInterrupt,omitempty" yaml:"allowInterrupt,omitempty"`

	// Delay between when the Lambda fulfillment function starts running and the start message is played. If the Lambda function returns before the delay is over, the start message isn't played.
	DelayInSeconds int `json:"delayInSeconds,omitempty" yaml:"delayInSeconds,omitempty"`

	// Between 1-5 configuration block message groups that contain start messages. Amazon Lex chooses one of the messages to play to the user. See `message_group`.
	MessageGroups []Lex_V2modelsIntentFulfillmentCodeHookFulfillmentUpdatesSpecificationStartResponseMessageGroup `json:"messageGroups,omitempty" yaml:"messageGroups,omitempty"`
}
