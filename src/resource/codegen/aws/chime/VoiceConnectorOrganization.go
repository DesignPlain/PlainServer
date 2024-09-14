package chime

import types "libds/aws/types"

type VoiceConnectorOrganization struct {
	// When origination settings are disabled, inbound calls are not enabled for your Amazon Chime Voice Connector.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	// Set of call distribution properties defined for your SIP hosts. See route below for more details. Minimum of 1. Maximum of 20.
	Routes []types.Chime_VoiceConnectorOrganizationRoute `json:"routes,omitempty" yaml:"routes,omitempty"`

	// The Amazon Chime Voice Connector ID.
	VoiceConnectorId string `json:"voiceConnectorId,omitempty" yaml:"voiceConnectorId,omitempty"`
}
