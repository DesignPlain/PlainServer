package pinpoint

type AdmChannel struct {
	// Client ID (part of OAuth Credentials) obtained via Amazon Developer Account.
	ClientId string `json:"clientId,omitempty" yaml:"clientId,omitempty"`

	// Client Secret (part of OAuth Credentials) obtained via Amazon Developer Account.
	ClientSecret string `json:"clientSecret,omitempty" yaml:"clientSecret,omitempty"`

	// Specifies whether to enable the channel. Defaults to `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The application ID.
	ApplicationId string `json:"applicationId,omitempty" yaml:"applicationId,omitempty"`
}
