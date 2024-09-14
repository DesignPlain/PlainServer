package types

type Compute_SecurityScanConfigAuthentication struct {
	/*
	   Describes authentication configuration that uses a custom account.
	   Structure is documented below.
	*/
	CustomAccount Compute_SecurityScanConfigAuthenticationCustomAccount `json:"customAccount,omitempty" yaml:"customAccount,omitempty"`

	/*
	   Describes authentication configuration that uses a Google account.
	   Structure is documented below.
	*/
	GoogleAccount Compute_SecurityScanConfigAuthenticationGoogleAccount `json:"googleAccount,omitempty" yaml:"googleAccount,omitempty"`
}
