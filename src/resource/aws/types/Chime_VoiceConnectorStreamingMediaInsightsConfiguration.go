package types

type Chime_VoiceConnectorStreamingMediaInsightsConfiguration struct {
	// The media insights configuration that will be invoked by the Voice Connector.
	ConfigurationArn string `json:"configurationArn,omitempty" yaml:"configurationArn,omitempty"`

	// When `true`, the media insights configuration is not enabled. Defaults to `false`.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
}
