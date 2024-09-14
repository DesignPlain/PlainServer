package waf

import types "libds/aws/types"

type Rule struct {
	// The name or description for the Amazon CloudWatch metric of this rule. The name can contain only alphanumeric characters (A-Z, a-z, 0-9); the name can't contain whitespace.
	MetricName string `json:"metricName,omitempty" yaml:"metricName,omitempty"`

	// The name or description of the rule.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The objects to include in a rule (documented below).
	Predicates []types.Waf_RulePredicate `json:"predicates,omitempty" yaml:"predicates,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
