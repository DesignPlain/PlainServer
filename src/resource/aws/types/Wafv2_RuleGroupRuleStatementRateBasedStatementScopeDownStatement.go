package types

type Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatement struct {
	// A rule statement that defines a string match search against labels that have been added to the web request by rules that have already run in the web ACL. See Label Match Statement below for details.
	LabelMatchStatement Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementLabelMatchStatement `json:"labelMatchStatement,omitempty" yaml:"labelMatchStatement,omitempty"`

	// A rule statement used to search web request components for matches with regular expressions. See Regex Pattern Set Reference Statement below for details.
	RegexPatternSetReferenceStatement Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatement `json:"regexPatternSetReferenceStatement,omitempty" yaml:"regexPatternSetReferenceStatement,omitempty"`

	// A logical rule statement used to combine other rule statements with AND logic. See AND Statement below for details.
	AndStatement Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementAndStatement `json:"andStatement,omitempty" yaml:"andStatement,omitempty"`

	// A rule statement that defines a string match search for AWS WAF to apply to web requests. See Byte Match Statement below for details.
	ByteMatchStatement Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementByteMatchStatement `json:"byteMatchStatement,omitempty" yaml:"byteMatchStatement,omitempty"`

	// A rule statement used to detect web requests coming from particular IP addresses or address ranges. See IP Set Reference Statement below for details.
	IpSetReferenceStatement Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementIpSetReferenceStatement `json:"ipSetReferenceStatement,omitempty" yaml:"ipSetReferenceStatement,omitempty"`

	// A rule statement used to search web request components for a match against a single regular expression. See Regex Match Statement below for details.
	RegexMatchStatement Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementRegexMatchStatement `json:"regexMatchStatement,omitempty" yaml:"regexMatchStatement,omitempty"`

	// A rule statement that compares a number of bytes against the size of a request component, using a comparison operator, such as greater than (>) or less than (<). See Size Constraint Statement below for more details.
	SizeConstraintStatement Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementSizeConstraintStatement `json:"sizeConstraintStatement,omitempty" yaml:"sizeConstraintStatement,omitempty"`

	// An SQL injection match condition identifies the part of web requests, such as the URI or the query string, that you want AWS WAF to inspect. See SQL Injection Match Statement below for details.
	SqliMatchStatement Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementSqliMatchStatement `json:"sqliMatchStatement,omitempty" yaml:"sqliMatchStatement,omitempty"`

	// A rule statement that defines a cross-site scripting (XSS) match search for AWS WAF to apply to web requests. See XSS Match Statement below for details.
	XssMatchStatement Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementXssMatchStatement `json:"xssMatchStatement,omitempty" yaml:"xssMatchStatement,omitempty"`

	// A rule statement used to identify web requests based on country of origin. See GEO Match Statement below for details.
	GeoMatchStatement Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementGeoMatchStatement `json:"geoMatchStatement,omitempty" yaml:"geoMatchStatement,omitempty"`

	// A logical rule statement used to negate the results of another rule statement. See NOT Statement below for details.
	NotStatement Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementNotStatement `json:"notStatement,omitempty" yaml:"notStatement,omitempty"`

	// A logical rule statement used to combine other rule statements with OR logic. See OR Statement below for details.
	OrStatement Wafv2_RuleGroupRuleStatementRateBasedStatementScopeDownStatementOrStatement `json:"orStatement,omitempty" yaml:"orStatement,omitempty"`
}
