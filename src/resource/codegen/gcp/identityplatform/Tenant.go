package identityplatform

type Tenant struct {
	// Whether to enable email link user authentication.
	EnableEmailLinkSignin bool `json:"enableEmailLinkSignin,omitempty" yaml:"enableEmailLinkSignin,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Whether to allow email/password user authentication.
	AllowPasswordSignup bool `json:"allowPasswordSignup,omitempty" yaml:"allowPasswordSignup,omitempty"`

	/*
	   Whether authentication is disabled for the tenant. If true, the users under
	   the disabled tenant are not allowed to sign-in. Admins of the disabled tenant
	   are not able to manage its users.
	*/
	DisableAuth bool `json:"disableAuth,omitempty" yaml:"disableAuth,omitempty"`

	/*
	   Human friendly display name of the tenant.


	   - - -
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
