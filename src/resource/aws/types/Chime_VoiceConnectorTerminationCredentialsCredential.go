package types

type Chime_VoiceConnectorTerminationCredentialsCredential struct {
	// RFC2617 compliant password associated with the SIP credentials.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// RFC2617 compliant username associated with the SIP credentials.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
