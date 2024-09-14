package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetRequestInspection struct {
	// The names of the fields in the request payload that contain your customer's primary physical address. See `address_fields` for more details.
	AddressFields Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetRequestInspectionAddressFields `json:"addressFields,omitempty" yaml:"addressFields,omitempty"`

	// The name of the field in the request payload that contains your customer's email. See `email_field` for more details.
	EmailField Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetRequestInspectionEmailField `json:"emailField,omitempty" yaml:"emailField,omitempty"`

	// Details about your login page password field. See `password_field` for more details.
	PasswordField Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetRequestInspectionPasswordField `json:"passwordField,omitempty" yaml:"passwordField,omitempty"`

	// The payload type for your login endpoint, either JSON or form encoded.
	PayloadType string `json:"payloadType,omitempty" yaml:"payloadType,omitempty"`

	// The names of the fields in the request payload that contain your customer's primary phone number. See `phone_number_fields` for more details.
	PhoneNumberFields Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetRequestInspectionPhoneNumberFields `json:"phoneNumberFields,omitempty" yaml:"phoneNumberFields,omitempty"`

	// Details about your login page username field. See `username_field` for more details.
	UsernameField Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetRequestInspectionUsernameField `json:"usernameField,omitempty" yaml:"usernameField,omitempty"`
}
