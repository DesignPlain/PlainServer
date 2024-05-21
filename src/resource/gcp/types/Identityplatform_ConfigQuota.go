package types

type Identityplatform_ConfigQuota struct {
	/*
	   Quota for the Signup endpoint, if overwritten. Signup quota is measured in sign ups per project per hour per IP.
	   Structure is documented below.
	*/
	SignUpQuotaConfig Identityplatform_ConfigQuotaSignUpQuotaConfig `json:"signUpQuotaConfig,omitempty" yaml:"signUpQuotaConfig,omitempty"`
}
