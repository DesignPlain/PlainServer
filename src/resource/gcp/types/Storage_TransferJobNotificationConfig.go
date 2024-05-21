package types

type Storage_TransferJobNotificationConfig struct {
	// Event types for which a notification is desired. If empty, send notifications for all event  The valid types are "TRANSFER_OPERATION_SUCCESS", "TRANSFER_OPERATION_FAILED", "TRANSFER_OPERATION_ABORTED".
	EventTypes []string `json:"eventTypes,omitempty" yaml:"eventTypes,omitempty"`

	// The desired format of the notification message payloads. One of "NONE" or "JSON".
	PayloadFormat string `json:"payloadFormat,omitempty" yaml:"payloadFormat,omitempty"`

	// The Topic.name of the Pub/Sub topic to which to publish notifications. Must be of the format: projects/{project}/topics/{topic}. Not matching this format results in an INVALID_ARGUMENT error.
	PubsubTopic string `json:"pubsubTopic,omitempty" yaml:"pubsubTopic,omitempty"`
}
