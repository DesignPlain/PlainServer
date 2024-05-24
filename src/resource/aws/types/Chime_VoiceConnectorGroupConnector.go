package types

type Chime_VoiceConnectorGroupConnector struct {
	// The priority associated with the Amazon Chime Voice Connector, with 1 being the highest priority. Higher priority Amazon Chime Voice Connectors are attempted first.
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// The Amazon Chime Voice Connector ID.
	VoiceConnectorId string `json:"voiceConnectorId,omitempty" yaml:"voiceConnectorId,omitempty"`
}
