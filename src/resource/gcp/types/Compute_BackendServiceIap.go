package types

type Compute_BackendServiceIap struct {
	/*
	   OAuth2 Client Secret for IAP
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Oauth2ClientSecret string `json:"oauth2ClientSecret,omitempty" yaml:"oauth2ClientSecret,omitempty"`

	/*
	   (Output)
	   OAuth2 Client Secret SHA-256 for IAP
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Oauth2ClientSecretSha256 string `json:"oauth2ClientSecretSha256,omitempty" yaml:"oauth2ClientSecretSha256,omitempty"`

	// OAuth2 Client ID for IAP
	Oauth2ClientId string `json:"oauth2ClientId,omitempty" yaml:"oauth2ClientId,omitempty"`
}
