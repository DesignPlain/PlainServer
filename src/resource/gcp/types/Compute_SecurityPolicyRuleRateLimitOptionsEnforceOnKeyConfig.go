package types

type Compute_SecurityPolicyRuleRateLimitOptionsEnforceOnKeyConfig struct {
	// Rate limit key name applicable only for the following key types:
	EnforceOnKeyName string `json:"enforceOnKeyName,omitempty" yaml:"enforceOnKeyName,omitempty"`

	// Determines the key to enforce the `rate_limit_threshold` on. If not specified, defaults to `ALL`.
	EnforceOnKeyType string `json:"enforceOnKeyType,omitempty" yaml:"enforceOnKeyType,omitempty"`
}
