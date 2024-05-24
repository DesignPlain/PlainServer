package chime

type VoiceConnectorTermination struct {
	// The Amazon Chime Voice Connector ID.
	VoiceConnectorId string `json:"voiceConnectorId,omitempty" yaml:"voiceConnectorId,omitempty"`

	// The countries to which calls are allowed, in ISO 3166-1 alpha-2 format.
	CallingRegions []string `json:"callingRegions,omitempty" yaml:"callingRegions,omitempty"`

	// The IP addresses allowed to make calls, in CIDR format.
	CidrAllowLists []string `json:"cidrAllowLists,omitempty" yaml:"cidrAllowLists,omitempty"`

	// The limit on calls per second. Max value based on account service quota. Default value of `1`.
	CpsLimit int `json:"cpsLimit,omitempty" yaml:"cpsLimit,omitempty"`

	// The default caller ID phone number.
	DefaultPhoneNumber string `json:"defaultPhoneNumber,omitempty" yaml:"defaultPhoneNumber,omitempty"`

	// When termination settings are disabled, outbound calls can not be made.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`
}
