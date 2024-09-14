package chime

import types "libds/aws/types"

type VoiceConnectorGroup struct {
	// The Amazon Chime Voice Connectors to route inbound calls to.
	Connectors []types.Chime_VoiceConnectorGroupConnector `json:"connectors,omitempty" yaml:"connectors,omitempty"`

	// The name of the Amazon Chime Voice Connector group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
