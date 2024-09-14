package types

type Networkfirewall_RuleGroupRuleGroupRulesSourceStatefulRuleRuleOption struct {
	/*
	   Keyword defined by open source detection systems like Snort or Suricata for stateful rule inspection.
	   See [Snort General Rule Options](http://manual-snort-org.s3-website-us-east-1.amazonaws.com/node31.html) or [Suricata Rule Options](https://suricata.readthedocs.io/en/suricata-5.0.1/rules/intro.html#rule-options) for more details.
	*/
	Keyword string `json:"keyword,omitempty" yaml:"keyword,omitempty"`

	// Set of strings for additional settings to use in stateful rule inspection.
	Settings []string `json:"settings,omitempty" yaml:"settings,omitempty"`
}
