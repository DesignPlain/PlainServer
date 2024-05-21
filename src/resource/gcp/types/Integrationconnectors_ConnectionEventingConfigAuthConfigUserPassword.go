package types

type Integrationconnectors_ConnectionEventingConfigAuthConfigUserPassword struct {
	// Username for Authentication.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	/*
	   Password for Authentication.
	   Structure is documented below.
	*/
	Password Integrationconnectors_ConnectionEventingConfigAuthConfigUserPasswordPassword `json:"password,omitempty" yaml:"password,omitempty"`
}
