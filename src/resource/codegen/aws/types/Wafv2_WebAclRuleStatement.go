package types

type Wafv2_WebAclRuleStatement struct {
	// Logical rule statement used to negate the results of another rule statement. See `not_statement` below for details.
	NotStatement Wafv2_WebAclRuleStatementNotStatement `json:"notStatement,omitempty" yaml:"notStatement,omitempty"`

	// Rule statement that compares a number of bytes against the size of a request component, using a comparison operator, such as greater than (>) or less than (<). See `size_constraint_statement` below for more details.
	SizeConstraintStatement Wafv2_WebAclRuleStatementSizeConstraintStatement `json:"sizeConstraintStatement,omitempty" yaml:"sizeConstraintStatement,omitempty"`

	// Rule statement used to identify web requests based on country of origin. See `geo_match_statement` below for details.
	GeoMatchStatement Wafv2_WebAclRuleStatementGeoMatchStatement `json:"geoMatchStatement,omitempty" yaml:"geoMatchStatement,omitempty"`

	// Rule statement used to detect web requests coming from particular IP addresses or address ranges. See `ip_set_reference_statement` below for details.
	IpSetReferenceStatement Wafv2_WebAclRuleStatementIpSetReferenceStatement `json:"ipSetReferenceStatement,omitempty" yaml:"ipSetReferenceStatement,omitempty"`

	// Logical rule statement used to combine other rule statements with OR logic. See `or_statement` below for details.
	OrStatement Wafv2_WebAclRuleStatementOrStatement `json:"orStatement,omitempty" yaml:"orStatement,omitempty"`

	// Rule statement used to search web request components for matches with regular expressions. See `regex_pattern_set_reference_statement` below for details.
	RegexPatternSetReferenceStatement Wafv2_WebAclRuleStatementRegexPatternSetReferenceStatement `json:"regexPatternSetReferenceStatement,omitempty" yaml:"regexPatternSetReferenceStatement,omitempty"`

	// Logical rule statement used to combine other rule statements with AND logic. See `and_statement` below for details.
	AndStatement Wafv2_WebAclRuleStatementAndStatement `json:"andStatement,omitempty" yaml:"andStatement,omitempty"`

	// Rule statement that defines a string match search against labels that have been added to the web request by rules that have already run in the web ACL. See `label_match_statement` below for details.
	LabelMatchStatement Wafv2_WebAclRuleStatementLabelMatchStatement `json:"labelMatchStatement,omitempty" yaml:"labelMatchStatement,omitempty"`

	// Rule statement used to search web request components for a match against a single regular expression. See `regex_match_statement` below for details.
	RegexMatchStatement Wafv2_WebAclRuleStatementRegexMatchStatement `json:"regexMatchStatement,omitempty" yaml:"regexMatchStatement,omitempty"`

	// Rule statement that defines a cross-site scripting (XSS) match search for AWS WAF to apply to web requests. See `xss_match_statement` below for details.
	XssMatchStatement Wafv2_WebAclRuleStatementXssMatchStatement `json:"xssMatchStatement,omitempty" yaml:"xssMatchStatement,omitempty"`

	// An SQL injection match condition identifies the part of web requests, such as the URI or the query string, that you want AWS WAF to inspect. See `sqli_match_statement` below for details.
	SqliMatchStatement Wafv2_WebAclRuleStatementSqliMatchStatement `json:"sqliMatchStatement,omitempty" yaml:"sqliMatchStatement,omitempty"`

	// Rule statement that defines a string match search for AWS WAF to apply to web requests. See `byte_match_statement` below for details.
	ByteMatchStatement Wafv2_WebAclRuleStatementByteMatchStatement `json:"byteMatchStatement,omitempty" yaml:"byteMatchStatement,omitempty"`

	// Rule statement used to run the rules that are defined in a managed rule group.  This statement can not be nested. See `managed_rule_group_statement` below for details.
	ManagedRuleGroupStatement Wafv2_WebAclRuleStatementManagedRuleGroupStatement `json:"managedRuleGroupStatement,omitempty" yaml:"managedRuleGroupStatement,omitempty"`

	// Rate-based rule tracks the rate of requests for each originating `IP address`, and triggers the rule action when the rate exceeds a limit that you specify on the number of requests in any `5-minute` time span. This statement can not be nested. See `rate_based_statement` below for details.
	RateBasedStatement Wafv2_WebAclRuleStatementRateBasedStatement `json:"rateBasedStatement,omitempty" yaml:"rateBasedStatement,omitempty"`

	// Rule statement used to run the rules that are defined in an WAFv2 Rule Group. See `rule_group_reference_statement` below for details.
	RuleGroupReferenceStatement Wafv2_WebAclRuleStatementRuleGroupReferenceStatement `json:"ruleGroupReferenceStatement,omitempty" yaml:"ruleGroupReferenceStatement,omitempty"`
}
