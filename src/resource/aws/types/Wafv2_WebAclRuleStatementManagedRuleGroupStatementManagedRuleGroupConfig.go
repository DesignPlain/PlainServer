package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfig struct {
	// The path of the login endpoint for your application.
	LoginPath string `json:"loginPath,omitempty" yaml:"loginPath,omitempty"`

	// Details about your login page password field. See `password_field` for more details.
	PasswordField Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigPasswordField `json:"passwordField,omitempty" yaml:"passwordField,omitempty"`

	// The payload type for your login endpoint, either JSON or form encoded.
	PayloadType string `json:"payloadType,omitempty" yaml:"payloadType,omitempty"`

	// Details about your login page username field. See `username_field` for more details.
	UsernameField Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigUsernameField `json:"usernameField,omitempty" yaml:"usernameField,omitempty"`

	// Additional configuration for using the Account Creation Fraud Prevention managed rule group. Use this to specify information such as the registration page of your application and the type of content to accept or reject from the client.
	AwsManagedRulesAcfpRuleSet Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSet `json:"awsManagedRulesAcfpRuleSet,omitempty" yaml:"awsManagedRulesAcfpRuleSet,omitempty"`

	// Additional configuration for using the Account Takeover Protection managed rule group. Use this to specify information such as the sign-in page of your application and the type of content to accept or reject from the client.
	AwsManagedRulesAtpRuleSet Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAtpRuleSet `json:"awsManagedRulesAtpRuleSet,omitempty" yaml:"awsManagedRulesAtpRuleSet,omitempty"`

	// Additional configuration for using the Bot Control managed rule group. Use this to specify the inspection level that you want to use. See `aws_managed_rules_bot_control_rule_set` for more details
	AwsManagedRulesBotControlRuleSet Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesBotControlRuleSet `json:"awsManagedRulesBotControlRuleSet,omitempty" yaml:"awsManagedRulesBotControlRuleSet,omitempty"`
}
