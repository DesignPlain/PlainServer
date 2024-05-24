package waf

import types "DesignSphere_Server/src/resource/aws/types"

type WebAcl struct {
	// Configuration blocks containing rules to associate with the web ACL and the settings for each rule. Detailed below.
	Rules []types.Waf_WebAclRule `json:"rules,omitempty" yaml:"rules,omitempty"`

	// Key-value map of resource tags. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configuration block with action that you want AWS WAF to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL. Detailed below.
	DefaultAction types.Waf_WebAclDefaultAction `json:"defaultAction,omitempty" yaml:"defaultAction,omitempty"`

	// Configuration block to enable WAF logging. Detailed below.
	LoggingConfiguration types.Waf_WebAclLoggingConfiguration `json:"loggingConfiguration,omitempty" yaml:"loggingConfiguration,omitempty"`

	// The name or description for the Amazon CloudWatch metric of this web ACL.
	MetricName string `json:"metricName,omitempty" yaml:"metricName,omitempty"`

	// The name or description of the web ACL.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
