package chime

import types "DesignSphere_Server/src/resource/aws/types"

type VoiceConnectorTerminationCredentials struct {
	// List of termination SIP credentials.
	Credentials []types.Chime_VoiceConnectorTerminationCredentialsCredential `json:"credentials,omitempty" yaml:"credentials,omitempty"`

	// Amazon Chime Voice Connector ID.
	VoiceConnectorId string `json:"voiceConnectorId,omitempty" yaml:"voiceConnectorId,omitempty"`
}
