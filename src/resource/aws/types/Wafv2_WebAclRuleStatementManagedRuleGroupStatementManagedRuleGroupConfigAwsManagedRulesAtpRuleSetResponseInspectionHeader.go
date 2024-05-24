package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAtpRuleSetResponseInspectionHeader struct {
	// Values in the response header with the specified name that indicate a failed login attempt.
	FailureValues []string `json:"failureValues,omitempty" yaml:"failureValues,omitempty"`

	// The name of the header to use.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Values in the response header with the specified name that indicate a successful login attempt.
	SuccessValues []string `json:"successValues,omitempty" yaml:"successValues,omitempty"`
}
