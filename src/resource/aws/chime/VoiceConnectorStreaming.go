package chime

import types "DesignSphere_Server/src/resource/aws/types"

type VoiceConnectorStreaming struct {
	// The media insights configuration. See `media_insights_configuration`.
	MediaInsightsConfiguration types.Chime_VoiceConnectorStreamingMediaInsightsConfiguration `json:"mediaInsightsConfiguration,omitempty" yaml:"mediaInsightsConfiguration,omitempty"`

	// The streaming notification targets. Valid Values: `EventBridge | SNS | SQS`
	StreamingNotificationTargets []string `json:"streamingNotificationTargets,omitempty" yaml:"streamingNotificationTargets,omitempty"`

	// The Amazon Chime Voice Connector ID.
	VoiceConnectorId string `json:"voiceConnectorId,omitempty" yaml:"voiceConnectorId,omitempty"`

	// The retention period, in hours, for the Amazon Kinesis data.
	DataRetention int `json:"dataRetention,omitempty" yaml:"dataRetention,omitempty"`

	// When true, media streaming to Amazon Kinesis is turned off. Default: `false`
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
}
