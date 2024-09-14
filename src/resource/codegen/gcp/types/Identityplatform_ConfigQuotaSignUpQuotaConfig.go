package types

type Identityplatform_ConfigQuotaSignUpQuotaConfig struct {
	// A sign up APIs quota that customers can override temporarily.
	Quota int `json:"quota,omitempty" yaml:"quota,omitempty"`

	// How long this quota will be active for. It is measurred in seconds, e.g., Example: "9.615s".
	QuotaDuration string `json:"quotaDuration,omitempty" yaml:"quotaDuration,omitempty"`

	// When this quota will take affect.
	StartTime string `json:"startTime,omitempty" yaml:"startTime,omitempty"`
}
