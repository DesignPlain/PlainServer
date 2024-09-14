package types

type Compute_SecurityPolicyRulePreconfiguredWafConfig struct {
	// An exclusion to apply during preconfigured WAF evaluation. Structure is documented below.
	Exclusions []Compute_SecurityPolicyRulePreconfiguredWafConfigExclusion `json:"exclusions,omitempty" yaml:"exclusions,omitempty"`
}
