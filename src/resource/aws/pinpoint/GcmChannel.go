package pinpoint

type GcmChannel struct {
	// The application ID.
	ApplicationId string `json:"applicationId,omitempty" yaml:"applicationId,omitempty"`

	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Platform credential API key from Google.
	ApiKey string `json:"apiKey,omitempty" yaml:"apiKey,omitempty"`
}
