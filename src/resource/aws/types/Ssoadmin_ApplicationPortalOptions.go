package types

type Ssoadmin_ApplicationPortalOptions struct {
	// Indicates whether this application is visible in the access portal. Valid values are `ENABLED` and `DISABLED`.
	Visibility string `json:"visibility,omitempty" yaml:"visibility,omitempty"`

	// Sign-in options for the access portal. See `sign_in_options` below.
	SignInOptions Ssoadmin_ApplicationPortalOptionsSignInOptions `json:"signInOptions,omitempty" yaml:"signInOptions,omitempty"`
}
