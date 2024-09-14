package types

type Integrationconnectors_ConnectionEventingConfigAuthConfig struct {
	/*
	   List containing additional auth configs.
	   Structure is documented below.
	*/
	AdditionalVariables []Integrationconnectors_ConnectionEventingConfigAuthConfigAdditionalVariable `json:"additionalVariables,omitempty" yaml:"additionalVariables,omitempty"`

	// The type of authentication configured.
	AuthKey string `json:"authKey,omitempty" yaml:"authKey,omitempty"`

	/*
	   authType of the Connection
	   Possible values are: `USER_PASSWORD`.
	*/
	AuthType string `json:"authType,omitempty" yaml:"authType,omitempty"`

	/*
	   User password for Authentication.
	   Structure is documented below.
	*/
	UserPassword Integrationconnectors_ConnectionEventingConfigAuthConfigUserPassword `json:"userPassword,omitempty" yaml:"userPassword,omitempty"`
}
