package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspectionStatusCode struct {
	// Status codes in the response that indicate a failed login attempt.
	FailureCodes []int `json:"failureCodes,omitempty" yaml:"failureCodes,omitempty"`

	// Status codes in the response that indicate a successful login attempt.
	SuccessCodes []int `json:"successCodes,omitempty" yaml:"successCodes,omitempty"`
}
