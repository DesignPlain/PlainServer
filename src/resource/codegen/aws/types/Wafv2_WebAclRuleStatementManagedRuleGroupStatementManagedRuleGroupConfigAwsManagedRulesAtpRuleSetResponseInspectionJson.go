package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAtpRuleSetResponseInspectionJson struct {
	// Values in the response header with the specified name that indicate a successful login attempt.
	SuccessValues []string `json:"successValues,omitempty" yaml:"successValues,omitempty"`

	// Values in the response header with the specified name that indicate a failed login attempt.
	FailureValues []string `json:"failureValues,omitempty" yaml:"failureValues,omitempty"`

	// The identifier for the value to match against in the JSON.
	Identifier string `json:"identifier,omitempty" yaml:"identifier,omitempty"`
}
