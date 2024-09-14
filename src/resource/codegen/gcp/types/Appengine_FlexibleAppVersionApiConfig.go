package types

type Appengine_FlexibleAppVersionApiConfig struct {
	/*
	   Security (HTTPS) enforcement for this URL.
	   Possible values are: `SECURE_DEFAULT`, `SECURE_NEVER`, `SECURE_OPTIONAL`, `SECURE_ALWAYS`.
	*/
	SecurityLevel string `json:"securityLevel,omitempty" yaml:"securityLevel,omitempty"`

	// URL to serve the endpoint at.
	Url string `json:"url,omitempty" yaml:"url,omitempty"`

	/*
	   Action to take when users access resources that require authentication.
	   Default value is `AUTH_FAIL_ACTION_REDIRECT`.
	   Possible values are: `AUTH_FAIL_ACTION_REDIRECT`, `AUTH_FAIL_ACTION_UNAUTHORIZED`.
	*/
	AuthFailAction string `json:"authFailAction,omitempty" yaml:"authFailAction,omitempty"`

	/*
	   Level of login required to access this resource.
	   Default value is `LOGIN_OPTIONAL`.
	   Possible values are: `LOGIN_OPTIONAL`, `LOGIN_ADMIN`, `LOGIN_REQUIRED`.
	*/
	Login string `json:"login,omitempty" yaml:"login,omitempty"`

	// Path to the script from the application root directory.
	Script string `json:"script,omitempty" yaml:"script,omitempty"`
}
