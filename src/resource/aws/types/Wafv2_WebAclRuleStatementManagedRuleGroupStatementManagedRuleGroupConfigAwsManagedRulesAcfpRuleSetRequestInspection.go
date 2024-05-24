package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetRequestInspection struct {
	//
	EmailField Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetRequestInspectionEmailField `json:"emailField,omitempty" yaml:"emailField,omitempty"`

	// Details about your login page password field. See `password_field` for more details.
	PasswordField Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetRequestInspectionPasswordField `json:"passwordField,omitempty" yaml:"passwordField,omitempty"`

	// The payload type for your login endpoint, either JSON or form encoded.
	PayloadType string `json:"payloadType,omitempty" yaml:"payloadType,omitempty"`

	// Details about your login page username field. See `username_field` for more details.
	UsernameField Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetRequestInspectionUsernameField `json:"usernameField,omitempty" yaml:"usernameField,omitempty"`
}
