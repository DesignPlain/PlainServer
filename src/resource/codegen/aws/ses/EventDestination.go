package ses

import types "libds/aws/types"

type EventDestination struct {
	// If true, the event destination will be enabled
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Send the events to a kinesis firehose destination
	KinesisDestination types.Ses_EventDestinationKinesisDestination `json:"kinesisDestination,omitempty" yaml:"kinesisDestination,omitempty"`

	// A list of matching types. May be any of `"send"`, `"reject"`, `"bounce"`, `"complaint"`, `"delivery"`, `"open"`, `"click"`, or `"renderingFailure"`.
	MatchingTypes []string `json:"matchingTypes,omitempty" yaml:"matchingTypes,omitempty"`

	// The name of the event destination
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Send the events to an SNS Topic destination

	   > --NOTE:-- You can specify `"cloudwatch_destination"` or `"kinesis_destination"` but not both
	*/
	SnsDestination types.Ses_EventDestinationSnsDestination `json:"snsDestination,omitempty" yaml:"snsDestination,omitempty"`

	// CloudWatch destination for the events
	CloudwatchDestinations []types.Ses_EventDestinationCloudwatchDestination `json:"cloudwatchDestinations,omitempty" yaml:"cloudwatchDestinations,omitempty"`

	// The name of the configuration set
	ConfigurationSetName string `json:"configurationSetName,omitempty" yaml:"configurationSetName,omitempty"`
}
