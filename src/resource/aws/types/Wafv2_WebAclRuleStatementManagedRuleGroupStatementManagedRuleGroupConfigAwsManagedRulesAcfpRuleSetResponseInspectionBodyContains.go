package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspectionBodyContains struct {
	// Strings in the body of the response that indicate a failed login attempt.
	FailureStrings []string `json:"failureStrings,omitempty" yaml:"failureStrings,omitempty"`

	// Strings in the body of the response that indicate a successful login attempt.
	SuccessStrings []string `json:"successStrings,omitempty" yaml:"successStrings,omitempty"`
}
