package types

type Integrationconnectors_ConnectionAuthConfig struct {
	/*
	   Parameters to support Oauth 2.0 Auth Code Grant Authentication.
	   Structure is documented below.
	*/
	Oauth2AuthCodeFlow Integrationconnectors_ConnectionAuthConfigOauth2AuthCodeFlow `json:"oauth2AuthCodeFlow,omitempty" yaml:"oauth2AuthCodeFlow,omitempty"`

	/*
	   OAuth3 Client Credentials for Authentication.
	   Structure is documented below.
	*/
	Oauth2ClientCredentials Integrationconnectors_ConnectionAuthConfigOauth2ClientCredentials `json:"oauth2ClientCredentials,omitempty" yaml:"oauth2ClientCredentials,omitempty"`

	/*
	   OAuth2 JWT Bearer for Authentication.
	   Structure is documented below.
	*/
	Oauth2JwtBearer Integrationconnectors_ConnectionAuthConfigOauth2JwtBearer `json:"oauth2JwtBearer,omitempty" yaml:"oauth2JwtBearer,omitempty"`

	/*
	   SSH Public Key for Authentication.
	   Structure is documented below.
	*/
	SshPublicKey Integrationconnectors_ConnectionAuthConfigSshPublicKey `json:"sshPublicKey,omitempty" yaml:"sshPublicKey,omitempty"`

	/*
	   User password for Authentication.
	   Structure is documented below.
	*/
	UserPassword Integrationconnectors_ConnectionAuthConfigUserPassword `json:"userPassword,omitempty" yaml:"userPassword,omitempty"`

	/*
	   List containing additional auth configs.
	   Structure is documented below.
	*/
	AdditionalVariables []Integrationconnectors_ConnectionAuthConfigAdditionalVariable `json:"additionalVariables,omitempty" yaml:"additionalVariables,omitempty"`

	// The type of authentication configured.
	AuthKey string `json:"authKey,omitempty" yaml:"authKey,omitempty"`

	/*
	   authType of the Connection
	   Possible values are: `USER_PASSWORD`.
	*/
	AuthType string `json:"authType,omitempty" yaml:"authType,omitempty"`
}
