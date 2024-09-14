package types

type Sesv2_ConfigurationSetEventDestinationEventDestination struct {
	// An object that defines an Amazon SNS destination for email events. See `sns_destination` Block for details.
	SnsDestination Sesv2_ConfigurationSetEventDestinationEventDestinationSnsDestination `json:"snsDestination,omitempty" yaml:"snsDestination,omitempty"`

	// An object that defines an Amazon CloudWatch destination for email events. See `cloud_watch_destination` Block for details.
	CloudWatchDestination Sesv2_ConfigurationSetEventDestinationEventDestinationCloudWatchDestination `json:"cloudWatchDestination,omitempty" yaml:"cloudWatchDestination,omitempty"`

	// When the event destination is enabled, the specified event types are sent to the destinations. Default: `false`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	EventBridgeDestination Sesv2_ConfigurationSetEventDestinationEventDestinationEventBridgeDestination `json:"eventBridgeDestination,omitempty" yaml:"eventBridgeDestination,omitempty"`

	// An object that defines an Amazon Kinesis Data Firehose destination for email events. See `kinesis_firehose_destination` Block for details.
	KinesisFirehoseDestination Sesv2_ConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestination `json:"kinesisFirehoseDestination,omitempty" yaml:"kinesisFirehoseDestination,omitempty"`

	// An array that specifies which events the Amazon SES API v2 should send to the destinations. Valid values: `SEND`, `REJECT`, `BOUNCE`, `COMPLAINT`, `DELIVERY`, `OPEN`, `CLICK`, `RENDERING_FAILURE`, `DELIVERY_DELAY`, `SUBSCRIPTION`.
	MatchingEventTypes []string `json:"matchingEventTypes,omitempty" yaml:"matchingEventTypes,omitempty"`

	// An object that defines an Amazon Pinpoint project destination for email events. See `pinpoint_destination` Block for details.
	PinpointDestination Sesv2_ConfigurationSetEventDestinationEventDestinationPinpointDestination `json:"pinpointDestination,omitempty" yaml:"pinpointDestination,omitempty"`
}
