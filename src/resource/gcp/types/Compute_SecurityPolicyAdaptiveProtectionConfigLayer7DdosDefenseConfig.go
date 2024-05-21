package types

type Compute_SecurityPolicyAdaptiveProtectionConfigLayer7DdosDefenseConfig struct {
	// If set to true, enables CAAP for L7 DDoS detection.
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`

	// Rule visibility can be one of the following:
	RuleVisibility string `json:"ruleVisibility,omitempty" yaml:"ruleVisibility,omitempty"`
}
