package types

type Compute_SecurityPolicyRuleRateLimitOptionsExceedRedirectOptions struct {
	// Target for the redirect action. This is required if the type is `EXTERNAL_302` and cannot be specified for `GOOGLE_RECAPTCHA`.
	Target string `json:"target,omitempty" yaml:"target,omitempty"`

	// Type of the redirect action.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
