package types

type Compute_SecurityScanConfigAuthenticationGoogleAccount struct {
	/*
	   The password of the Google account. The credential is stored encrypted
	   in GCP.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// The user name of the Google account.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
