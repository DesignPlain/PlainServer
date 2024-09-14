package networkfirewall

import types "libds/aws/types"

type RuleGroup struct {
	// Whether the rule group is stateless (containing stateless rules) or stateful (containing stateful rules). Valid values include: `STATEFUL` or `STATELESS`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The maximum number of operating resources that this rule group can use. For a stateless rule group, the capacity required is the sum of the capacity requirements of the individual rules. For a stateful rule group, the minimum capacity required is the number of individual rules.
	Capacity int `json:"capacity,omitempty" yaml:"capacity,omitempty"`

	// A friendly description of the rule group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// KMS encryption configuration settings. See Encryption Configuration below for details.
	EncryptionConfiguration types.Networkfirewall_RuleGroupEncryptionConfiguration `json:"encryptionConfiguration,omitempty" yaml:"encryptionConfiguration,omitempty"`

	// A friendly name of the rule group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A configuration block that defines the rule group rules. Required unless `rules` is specified. See Rule Group below for details.
	RuleGroup types.Networkfirewall_RuleGroupRuleGroup `json:"ruleGroup,omitempty" yaml:"ruleGroup,omitempty"`

	// The stateful rule group rules specifications in Suricata file format, with one rule per line. Use this to import your existing Suricata compatible rule groups. Required unless `rule_group` is specified.
	Rules string `json:"rules,omitempty" yaml:"rules,omitempty"`

	// A map of key:value pairs to associate with the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
