package chime

type VoiceConnectorLogging struct {
	// When true, enables logging of detailed media metrics for Voice Connectors to Amazon CloudWatch logs.
	EnableMediaMetricLogs bool `json:"enableMediaMetricLogs,omitempty" yaml:"enableMediaMetricLogs,omitempty"`

	// When true, enables SIP message logs for sending to Amazon CloudWatch Logs.
	EnableSipLogs bool `json:"enableSipLogs,omitempty" yaml:"enableSipLogs,omitempty"`

	// The Amazon Chime Voice Connector ID.
	VoiceConnectorId string `json:"voiceConnectorId,omitempty" yaml:"voiceConnectorId,omitempty"`
}
