package types

type Cognito_ManagedUserPoolClientTokenValidityUnits struct {
	// Time unit for the value in `refresh_token_validity` and defaults to `days`.
	RefreshToken string `json:"refreshToken,omitempty" yaml:"refreshToken,omitempty"`

	// Time unit for the value in `access_token_validity` and defaults to `hours`.
	AccessToken string `json:"accessToken,omitempty" yaml:"accessToken,omitempty"`

	// Time unit for the value in `id_token_validity`, and it defaults to `hours`.
	IdToken string `json:"idToken,omitempty" yaml:"idToken,omitempty"`
}
