package chime

import types "libds/aws/types"

type SdkvoiceGlobalSettings struct {
	// The Voice Connector settings. See voice_connector.
	VoiceConnector types.Chime_SdkvoiceGlobalSettingsVoiceConnector `json:"voiceConnector,omitempty" yaml:"voiceConnector,omitempty"`
}
