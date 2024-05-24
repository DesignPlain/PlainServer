package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspection struct {
	// Configures inspection of the response body. See `body_contains` for more details.
	BodyContains Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspectionBodyContains `json:"bodyContains,omitempty" yaml:"bodyContains,omitempty"`

	// Configures inspection of the response header.See `header` for more details.
	Header Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspectionHeader `json:"header,omitempty" yaml:"header,omitempty"`

	// Configures inspection of the response JSON. See `json` for more details.
	Json Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspectionJson `json:"json,omitempty" yaml:"json,omitempty"`

	// Configures inspection of the response status code.See `status_code` for more details.
	StatusCode Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspectionStatusCode `json:"statusCode,omitempty" yaml:"statusCode,omitempty"`
}
