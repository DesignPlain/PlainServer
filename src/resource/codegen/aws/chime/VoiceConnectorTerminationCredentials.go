package chime

import types "libds/aws/types"

type VoiceConnectorTerminationCredentials struct {
	// Amazon Chime Voice Connector ID.
	VoiceConnectorId string `json:"voiceConnectorId,omitempty" yaml:"voiceConnectorId,omitempty"`

	// List of termination SIP credentials.
	Credentials []types.Chime_VoiceConnectorTerminationCredentialsCredential `json:"credentials,omitempty" yaml:"credentials,omitempty"`
}
