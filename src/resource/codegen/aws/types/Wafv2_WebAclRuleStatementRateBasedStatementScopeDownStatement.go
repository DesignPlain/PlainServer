package types

type Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatement struct {
	// Rule statement used to identify web requests based on country of origin. See `geo_match_statement` below for details.
	GeoMatchStatement Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementGeoMatchStatement `json:"geoMatchStatement,omitempty" yaml:"geoMatchStatement,omitempty"`

	// Rule statement used to detect web requests coming from particular IP addresses or address ranges. See `ip_set_reference_statement` below for details.
	IpSetReferenceStatement Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementIpSetReferenceStatement `json:"ipSetReferenceStatement,omitempty" yaml:"ipSetReferenceStatement,omitempty"`

	// Rule statement that defines a string match search against labels that have been added to the web request by rules that have already run in the web ACL. See `label_match_statement` below for details.
	LabelMatchStatement Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementLabelMatchStatement `json:"labelMatchStatement,omitempty" yaml:"labelMatchStatement,omitempty"`

	// Logical rule statement used to combine other rule statements with OR logic. See `or_statement` below for details.
	OrStatement Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementOrStatement `json:"orStatement,omitempty" yaml:"orStatement,omitempty"`

	// Rule statement used to search web request components for a match against a single regular expression. See `regex_match_statement` below for details.
	RegexMatchStatement Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatement `json:"regexMatchStatement,omitempty" yaml:"regexMatchStatement,omitempty"`

	// Rule statement used to search web request components for matches with regular expressions. See `regex_pattern_set_reference_statement` below for details.
	RegexPatternSetReferenceStatement Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatement `json:"regexPatternSetReferenceStatement,omitempty" yaml:"regexPatternSetReferenceStatement,omitempty"`

	// Rule statement that compares a number of bytes against the size of a request component, using a comparison operator, such as greater than (>) or less than (<). See `size_constraint_statement` below for more details.
	SizeConstraintStatement Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementSizeConstraintStatement `json:"sizeConstraintStatement,omitempty" yaml:"sizeConstraintStatement,omitempty"`

	// Rule statement that defines a string match search for AWS WAF to apply to web requests. See `byte_match_statement` below for details.
	ByteMatchStatement Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementByteMatchStatement `json:"byteMatchStatement,omitempty" yaml:"byteMatchStatement,omitempty"`

	// Rule statement that defines a cross-site scripting (XSS) match search for AWS WAF to apply to web requests. See `xss_match_statement` below for details.
	XssMatchStatement Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementXssMatchStatement `json:"xssMatchStatement,omitempty" yaml:"xssMatchStatement,omitempty"`

	// An SQL injection match condition identifies the part of web requests, such as the URI or the query string, that you want AWS WAF to inspect. See `sqli_match_statement` below for details.
	SqliMatchStatement Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatement `json:"sqliMatchStatement,omitempty" yaml:"sqliMatchStatement,omitempty"`

	// Logical rule statement used to negate the results of another rule statement. See `not_statement` below for details.
	NotStatement Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementNotStatement `json:"notStatement,omitempty" yaml:"notStatement,omitempty"`

	// Logical rule statement used to combine other rule statements with AND logic. See `and_statement` below for details.
	AndStatement Wafv2_WebAclRuleStatementRateBasedStatementScopeDownStatementAndStatement `json:"andStatement,omitempty" yaml:"andStatement,omitempty"`
}
