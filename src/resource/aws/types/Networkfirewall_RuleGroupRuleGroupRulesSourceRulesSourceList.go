package types

type Networkfirewall_RuleGroupRuleGroupRulesSourceRulesSourceList struct {
	// String value to specify whether domains in the target list are allowed or denied access. Valid values: `ALLOWLIST`, `DENYLIST`.
	GeneratedRulesType string `json:"generatedRulesType,omitempty" yaml:"generatedRulesType,omitempty"`

	// Set of types of domain specifications that are provided in the `targets` argument. Valid values: `HTTP_HOST`, `TLS_SNI`.
	TargetTypes []string `json:"targetTypes,omitempty" yaml:"targetTypes,omitempty"`

	// Set of domains that you want to inspect for in your traffic flows.
	Targets []string `json:"targets,omitempty" yaml:"targets,omitempty"`
}
