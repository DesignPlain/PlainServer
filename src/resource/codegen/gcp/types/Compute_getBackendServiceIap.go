package types

type Compute_getBackendServiceIap struct {
	// OAuth2 Client ID for IAP
	Oauth2ClientId string `json:"oauth2ClientId,omitempty" yaml:"oauth2ClientId,omitempty"`

	// OAuth2 Client Secret for IAP
	Oauth2ClientSecret string `json:"oauth2ClientSecret,omitempty" yaml:"oauth2ClientSecret,omitempty"`

	// OAuth2 Client Secret SHA-256 for IAP
	Oauth2ClientSecretSha256 string `json:"oauth2ClientSecretSha256,omitempty" yaml:"oauth2ClientSecretSha256,omitempty"`
}
