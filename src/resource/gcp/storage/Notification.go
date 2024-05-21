package storage

type Notification struct {
	// A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
	CustomAttributes map[string]string `json:"customAttributes,omitempty" yaml:"customAttributes,omitempty"`

	// List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
	EventTypes []string `json:"eventTypes,omitempty" yaml:"eventTypes,omitempty"`

	// Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
	ObjectNamePrefix string `json:"objectNamePrefix,omitempty" yaml:"objectNamePrefix,omitempty"`

	// The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
	PayloadFormat string `json:"payloadFormat,omitempty" yaml:"payloadFormat,omitempty"`

	/*
	   The Cloud PubSub topic to which this subscription publishes. Expects either the
	   topic name, assumed to belong to the default GCP provider project, or the project-level name,
	   i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`. If the project is not set in the provider,
	   you will need to use the project-level name.

	   - - -
	*/
	Topic string `json:"topic,omitempty" yaml:"topic,omitempty"`

	// The name of the bucket.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`
}
