package wafv2

import types "libds/aws/types"

type RuleGroup struct {
	// A friendly name of the rule group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
	VisibilityConfig types.Wafv2_RuleGroupVisibilityConfig `json:"visibilityConfig,omitempty" yaml:"visibilityConfig,omitempty"`

	// The web ACL capacity units (WCUs) required for this rule group. See [here](https://docs.aws.amazon.com/waf/latest/APIReference/API_CreateRuleGroup.html#API_CreateRuleGroup_RequestSyntax) for general information and [here](https://docs.aws.amazon.com/waf/latest/developerguide/waf-rule-statements-list.html) for capacity specific information.
	Capacity int `json:"capacity,omitempty" yaml:"capacity,omitempty"`

	// Defines custom response bodies that can be referenced by `custom_response` actions. See Custom Response Body below for details.
	CustomResponseBodies []types.Wafv2_RuleGroupCustomResponseBody `json:"customResponseBodies,omitempty" yaml:"customResponseBodies,omitempty"`

	// A friendly description of the rule group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See Rules below for details.
	Rules []types.Wafv2_RuleGroupRule `json:"rules,omitempty" yaml:"rules,omitempty"`

	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`

	// An array of key:value pairs to associate with the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
