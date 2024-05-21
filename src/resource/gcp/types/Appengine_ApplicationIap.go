package types

type Appengine_ApplicationIap struct {
	/*
	   OAuth2 client secret to use for the authentication flow.
	   The SHA-256 hash of the value is returned in the oauth2ClientSecretSha256 field.
	*/
	Oauth2ClientSecret string `json:"oauth2ClientSecret,omitempty" yaml:"oauth2ClientSecret,omitempty"`

	// Hex-encoded SHA-256 hash of the client secret.
	Oauth2ClientSecretSha256 string `json:"oauth2ClientSecretSha256,omitempty" yaml:"oauth2ClientSecretSha256,omitempty"`

	/*
	   (Optional) Whether the serving infrastructure will authenticate and authorize all incoming requests.
	   (default is false)
	*/
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// OAuth2 client ID to use for the authentication flow.
	Oauth2ClientId string `json:"oauth2ClientId,omitempty" yaml:"oauth2ClientId,omitempty"`
}
