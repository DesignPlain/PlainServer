package types

type Identityplatform_ConfigClientPermissions struct {
	// When true, end users cannot delete their account on the associated project through any of our API methods
	DisabledUserDeletion bool `json:"disabledUserDeletion,omitempty" yaml:"disabledUserDeletion,omitempty"`

	// When true, end users cannot sign up for a new account on the associated project through any of our API methods
	DisabledUserSignup bool `json:"disabledUserSignup,omitempty" yaml:"disabledUserSignup,omitempty"`
}
