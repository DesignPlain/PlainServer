package wafv2

import types "libds/aws/types"

type WebAcl struct {
	// Specifies how AWS WAF should handle CAPTCHA evaluations on the ACL level (used by [AWS Bot Control](https://docs.aws.amazon.com/waf/latest/developerguide/aws-managed-rule-groups-bot.html)). See `captcha_config` below for details.
	CaptchaConfig types.Wafv2_WebAclCaptchaConfig `json:"captchaConfig,omitempty" yaml:"captchaConfig,omitempty"`

	// Action to perform if none of the `rules` contained in the WebACL match. See `default_action` below for details.
	DefaultAction types.Wafv2_WebAclDefaultAction `json:"defaultAction,omitempty" yaml:"defaultAction,omitempty"`

	// Rule blocks used to identify the web requests that you want to `allow`, `block`, or `count`. See `rule` below for details.
	Rules []types.Wafv2_WebAclRule `json:"rules,omitempty" yaml:"rules,omitempty"`

	// Defines and enables Amazon CloudWatch metrics and web request sample collection. See `visibility_config` below for details.
	VisibilityConfig types.Wafv2_WebAclVisibilityConfig `json:"visibilityConfig,omitempty" yaml:"visibilityConfig,omitempty"`

	// Map of key-value pairs to associate with the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies custom configurations for the associations between the web ACL and protected resources. See `association_config` below for details.
	AssociationConfig types.Wafv2_WebAclAssociationConfig `json:"associationConfig,omitempty" yaml:"associationConfig,omitempty"`

	// Specifies how AWS WAF should handle Challenge evaluations on the ACL level (used by [AWS Bot Control](https://docs.aws.amazon.com/waf/latest/developerguide/aws-managed-rule-groups-bot.html)). See `challenge_config` below for details.
	ChallengeConfig types.Wafv2_WebAclChallengeConfig `json:"challengeConfig,omitempty" yaml:"challengeConfig,omitempty"`

	// Defines custom response bodies that can be referenced by `custom_response` actions. See `custom_response_body` below for details.
	CustomResponseBodies []types.Wafv2_WebAclCustomResponseBody `json:"customResponseBodies,omitempty" yaml:"customResponseBodies,omitempty"`

	// Friendly description of the WebACL.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Friendly name of the WebACL.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Raw JSON string to allow more than three nested statements. Conflicts with `rule` attribute. This is for advanced use cases where more than 3 levels of nested statements are required. --There is no drift detection at this time--. If you use this attribute instead of `rule`, you will be foregoing drift detection. See the AWS [documentation](https://docs.aws.amazon.com/waf/latest/APIReference/API_CreateWebACL.html) for the JSON structure.
	RuleJson string `json:"ruleJson,omitempty" yaml:"ruleJson,omitempty"`

	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`

	// Specifies the domains that AWS WAF should accept in a web request token. This enables the use of tokens across multiple protected websites. When AWS WAF provides a token, it uses the domain of the AWS resource that the web ACL is protecting. If you don't specify a list of token domains, AWS WAF accepts tokens only for the domain of the protected resource. With a token domain list, AWS WAF accepts the resource's host domain plus all domains in the token domain list, including their prefixed subdomains.
	TokenDomains []string `json:"tokenDomains,omitempty" yaml:"tokenDomains,omitempty"`
}
