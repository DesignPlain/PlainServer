package wafregional

import types "libds/aws/types"

type WebAcl struct {
	// Set of configuration blocks containing rules for the web ACL. Detailed below.
	Rules []types.Wafregional_WebAclRule `json:"rules,omitempty" yaml:"rules,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
	DefaultAction types.Wafregional_WebAclDefaultAction `json:"defaultAction,omitempty" yaml:"defaultAction,omitempty"`

	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration types.Wafregional_WebAclLoggingConfiguration `json:"loggingConfiguration,omitempty" yaml:"loggingConfiguration,omitempty"`

	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName string `json:"metricName,omitempty" yaml:"metricName,omitempty"`

	// The name or description of the web ACL.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
