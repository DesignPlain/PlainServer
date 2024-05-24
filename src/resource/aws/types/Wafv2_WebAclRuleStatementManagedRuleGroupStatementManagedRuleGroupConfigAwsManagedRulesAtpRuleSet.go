package types

type Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAtpRuleSet struct {
	// Whether or not to allow the use of regular expressions in the login page path.
	EnableRegexInPath bool `json:"enableRegexInPath,omitempty" yaml:"enableRegexInPath,omitempty"`

	// The path of the login endpoint for your application.
	LoginPath string `json:"loginPath,omitempty" yaml:"loginPath,omitempty"`

	// The criteria for inspecting login requests, used by the ATP rule group to validate credentials usage. See `request_inspection` for more details.
	RequestInspection Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAtpRuleSetRequestInspection `json:"requestInspection,omitempty" yaml:"requestInspection,omitempty"`

	// The criteria for inspecting responses to login requests, used by the ATP rule group to track login failure rates. Note that Response Inspection is available only on web ACLs that protect CloudFront distributions. See `response_inspection` for more details.
	ResponseInspection Wafv2_WebAclRuleStatementManagedRuleGroupStatementManagedRuleGroupConfigAwsManagedRulesAtpRuleSetResponseInspection `json:"responseInspection,omitempty" yaml:"responseInspection,omitempty"`
}
