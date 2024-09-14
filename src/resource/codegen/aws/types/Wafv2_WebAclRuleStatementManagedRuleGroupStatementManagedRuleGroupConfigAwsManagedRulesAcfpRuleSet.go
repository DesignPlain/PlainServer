package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSet struct {
	// The path of the account creation endpoint for your application. This is the page on your website that accepts the completed registration form for a new user. This page must accept POST requests.
	CreationPath string `json:"creationPath,omitempty" yaml:"creationPath,omitempty"`

	// Whether or not to allow the use of regular expressions in the login page path.
	EnableRegexInPath bool `json:"enableRegexInPath,omitempty" yaml:"enableRegexInPath,omitempty"`

	// The path of the account registration endpoint for your application. This is the page on your website that presents the registration form to new users. This page must accept GET text/html requests.
	RegistrationPagePath string `json:"registrationPagePath,omitempty" yaml:"registrationPagePath,omitempty"`

	// The criteria for inspecting login requests, used by the ATP rule group to validate credentials usage. See `request_inspection` for more details.
	RequestInspection Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetRequestInspection `json:"requestInspection,omitempty" yaml:"requestInspection,omitempty"`

	// The criteria for inspecting responses to login requests, used by the ATP rule group to track login failure rates. Note that Response Inspection is available only on web ACLs that protect CloudFront distributions. See `response_inspection` for more details.
	ResponseInspection Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAcfpRuleSetResponseInspection `json:"responseInspection,omitempty" yaml:"responseInspection,omitempty"`
}
