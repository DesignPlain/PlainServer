package types

type Compute_SecurityPolicyRule struct {
	/*
	   When set to true, the `action` specified above is not enforced.
	   Stackdriver logs for requests that trigger a preview action are annotated as such.
	*/
	Preview bool `json:"preview,omitempty" yaml:"preview,omitempty"`

	// Can be specified if the `action` is `redirect`. Cannot be specified for other actions. Structure is documented below.
	RedirectOptions Compute_SecurityPolicyRuleRedirectOptions `json:"redirectOptions,omitempty" yaml:"redirectOptions,omitempty"`

	// Action to take when `match` matches the request. Valid values:
	Action string `json:"action,omitempty" yaml:"action,omitempty"`

	// An optional description of this rule. Max size is 64.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Additional actions that are performed on headers. Structure is documented below.
	HeaderAction Compute_SecurityPolicyRuleHeaderAction `json:"headerAction,omitempty" yaml:"headerAction,omitempty"`

	/*
	   A match condition that incoming traffic is evaluated against.
	   If it evaluates to true, the corresponding `action` is enforced. Structure is documented below.
	*/
	Match Compute_SecurityPolicyRuleMatch `json:"match,omitempty" yaml:"match,omitempty"`

	// Preconfigured WAF configuration to be applied for the rule. If the rule does not evaluate preconfigured WAF rules, i.e., if `evaluatePreconfiguredWaf()` is not used, this field will have no effect. Structure is documented below.
	PreconfiguredWafConfig Compute_SecurityPolicyRulePreconfiguredWafConfig `json:"preconfiguredWafConfig,omitempty" yaml:"preconfiguredWafConfig,omitempty"`

	/*
	   An unique positive integer indicating the priority of evaluation for a rule.
	   Rules are evaluated from highest priority (lowest numerically) to lowest priority (highest numerically) in order.
	*/
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	// Must be specified if the `action` is `rate_based_ban` or `throttle`. Cannot be specified for other actions. Structure is documented below.
	RateLimitOptions Compute_SecurityPolicyRuleRateLimitOptions `json:"rateLimitOptions,omitempty" yaml:"rateLimitOptions,omitempty"`
}
