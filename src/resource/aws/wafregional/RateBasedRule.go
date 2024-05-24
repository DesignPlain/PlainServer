package wafregional

import types "DesignSphere_Server/src/resource/aws/types"

type RateBasedRule struct {
	// Valid value is IP.
	RateKey string `json:"rateKey,omitempty" yaml:"rateKey,omitempty"`

	// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
	RateLimit int `json:"rateLimit,omitempty" yaml:"rateLimit,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName string `json:"metricName,omitempty" yaml:"metricName,omitempty"`

	// The name or description of the rule.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The objects to include in a rule (documented below).
	Predicates []types.Wafregional_RateBasedRulePredicate `json:"predicates,omitempty" yaml:"predicates,omitempty"`
}
