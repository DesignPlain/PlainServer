package types

type Waf_RulePredicate struct {
	// A unique identifier for a predicate in the rule, such as Byte Match Set ID or IPSet ID.
	DataId string `json:"dataId,omitempty" yaml:"dataId,omitempty"`

	/*
	   Set this to `false` if you want to allow, block, or count requests
	   based on the settings in the specified waf_byte_match_set, waf_ipset, aws_waf_size_constraint_set, aws.waf.SqlInjectionMatchSet or aws_waf_xss_match_set.
	   For example, if an IPSet includes the IP address `192.0.2.44`, AWS WAF will allow or block requests based on that IP address.
	   If set to `true`, AWS WAF will allow, block, or count requests based on all IP addresses except `192.0.2.44`.
	*/
	Negated bool `json:"negated,omitempty" yaml:"negated,omitempty"`

	// The type of predicate in a rule. Valid values: `ByteMatch`, `GeoMatch`, `IPMatch`, `RegexMatch`, `SizeConstraint`, `SqlInjectionMatch`, or `XssMatch`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
