package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAtpRuleSetResponseInspection struct {
	// Configures inspection of the response body. See `body_contains` for more details.
	BodyContains Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAtpRuleSetResponseInspectionBodyContains `json:"bodyContains,omitempty" yaml:"bodyContains,omitempty"`

	// Configures inspection of the response header.See `header` for more details.
	Header Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAtpRuleSetResponseInspectionHeader `json:"header,omitempty" yaml:"header,omitempty"`

	// Configures inspection of the response JSON. See `json` for more details.
	Json Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAtpRuleSetResponseInspectionJson `json:"json,omitempty" yaml:"json,omitempty"`

	// Configures inspection of the response status code.See `status_code` for more details.
	StatusCode Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAtpRuleSetResponseInspectionStatusCode `json:"statusCode,omitempty" yaml:"statusCode,omitempty"`
}
