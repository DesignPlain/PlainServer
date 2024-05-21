package types

type Integrationconnectors_ConnectionAuthConfigUserPassword struct {
	/*
	   Password for Authentication.
	   Structure is documented below.
	*/
	Password Integrationconnectors_ConnectionAuthConfigUserPasswordPassword `json:"password,omitempty" yaml:"password,omitempty"`

	// Username for Authentication.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
