package types

type Integrationconnectors_ConnectionEventingConfigAuthConfigUserPassword struct {
	/*
	   Password for Authentication.
	   Structure is documented below.
	*/
	Password Integrationconnectors_ConnectionEventingConfigAuthConfigUserPasswordPassword `json:"password,omitempty" yaml:"password,omitempty"`

	// Username for Authentication.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
