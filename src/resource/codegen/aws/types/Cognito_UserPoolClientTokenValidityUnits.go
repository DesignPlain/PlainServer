package types

type Cognito_UserPoolClientTokenValidityUnits struct {
	// Time unit in for the value in `access_token_validity`, defaults to `hours`.
	AccessToken string `json:"accessToken,omitempty" yaml:"accessToken,omitempty"`

	// Time unit in for the value in `id_token_validity`, defaults to `hours`.
	IdToken string `json:"idToken,omitempty" yaml:"idToken,omitempty"`

	// Time unit in for the value in `refresh_token_validity`, defaults to `days`.
	RefreshToken string `json:"refreshToken,omitempty" yaml:"refreshToken,omitempty"`
}
