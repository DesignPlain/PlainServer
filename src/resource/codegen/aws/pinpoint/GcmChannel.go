package pinpoint

type GcmChannel struct {
	//
	DefaultAuthenticationMethod string `json:"defaultAuthenticationMethod,omitempty" yaml:"defaultAuthenticationMethod,omitempty"`

	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	ServiceJson string `json:"serviceJson,omitempty" yaml:"serviceJson,omitempty"`

	// Platform credential API key from Google.
	ApiKey string `json:"apiKey,omitempty" yaml:"apiKey,omitempty"`

	// The application ID.
	ApplicationId string `json:"applicationId,omitempty" yaml:"applicationId,omitempty"`
}
