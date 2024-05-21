package types

type Compute_SecurityScanConfigAuthenticationCustomAccount struct {
	// The login form URL of the website.
	LoginUrl string `json:"loginUrl,omitempty" yaml:"loginUrl,omitempty"`

	/*
	   The password of the custom account. The credential is stored encrypted
	   in GCP.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// The user name of the custom account.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
