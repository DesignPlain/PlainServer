package wafregional

import types "DesignSphere_Server/src/resource/aws/types"

type RuleGroup struct {
	// A friendly name of the rule group
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A list of activated rules, see below
	ActivatedRules []types.Wafregional_RuleGroupActivatedRule `json:"activatedRules,omitempty" yaml:"activatedRules,omitempty"`

	// A friendly name for the metrics from the rule group
	MetricName string `json:"metricName,omitempty" yaml:"metricName,omitempty"`
}
