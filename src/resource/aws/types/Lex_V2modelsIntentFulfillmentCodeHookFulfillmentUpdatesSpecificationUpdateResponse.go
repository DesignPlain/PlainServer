package types

type Lex_V2modelsIntentFulfillmentCodeHookFulfillmentUpdatesSpecificationUpdateResponse struct {
	// Whether the user can interrupt the start message while it is playing.
	AllowInterrupt bool `json:"allowInterrupt,omitempty" yaml:"allowInterrupt,omitempty"`

	// Frequency that a message is sent to the user. When the period ends, Amazon Lex chooses a message from the message groups and plays it to the user. If the fulfillment Lambda returns before the first period ends, an update message is not played to the user.
	FrequencyInSeconds int `json:"frequencyInSeconds,omitempty" yaml:"frequencyInSeconds,omitempty"`

	// Between 1-5 configuration block message groups that contain start messages. Amazon Lex chooses one of the messages to play to the user. See `message_group`.
	MessageGroups []Lex_V2modelsIntentFulfillmentCodeHookFulfillmentUpdatesSpecificationUpdateResponseMessageGroup `json:"messageGroups,omitempty" yaml:"messageGroups,omitempty"`
}
