package cloudfront

import types "DesignSphere_Server/src/resource/aws/types"

type Distribution struct {
	// One or more origins for this distribution (multiples allowed).
	Origins []types.Cloudfront_DistributionOrigin `json:"origins,omitempty" yaml:"origins,omitempty"`

	// The restriction configuration for this distribution (maximum one).
	Restrictions types.Cloudfront_DistributionRestrictions `json:"restrictions,omitempty" yaml:"restrictions,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Identifier of a continuous deployment policy. This argument should only be set on a production distribution. See the `aws.cloudfront.ContinuousDeploymentPolicy` resource for additional details.
	ContinuousDeploymentPolicyId string `json:"continuousDeploymentPolicyId,omitempty" yaml:"continuousDeploymentPolicyId,omitempty"`

	// Default cache behavior for this distribution (maximum one). Requires either `cache_policy_id` (preferred) or `forwarded_values` (deprecated) be set.
	DefaultCacheBehavior types.Cloudfront_DistributionDefaultCacheBehavior `json:"defaultCacheBehavior,omitempty" yaml:"defaultCacheBehavior,omitempty"`

	// The logging configuration that controls how logs are written to your distribution (maximum one).
	LoggingConfig types.Cloudfront_DistributionLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`

	// The SSL configuration for this distribution (maximum one).
	ViewerCertificate types.Cloudfront_DistributionViewerCertificate `json:"viewerCertificate,omitempty" yaml:"viewerCertificate,omitempty"`

	// If enabled, the resource will wait for the distribution status to change from `InProgress` to `Deployed`. Setting this to`false` will skip the process. Default: `true`.
	WaitForDeployment bool `json:"waitForDeployment,omitempty" yaml:"waitForDeployment,omitempty"`

	// Unique identifier that specifies the AWS WAF web ACL, if any, to associate with this distribution. To specify a web ACL created using the latest version of AWS WAF (WAFv2), use the ACL ARN, for example `aws_wafv2_web_acl.example.arn`. To specify a web ACL created using AWS WAF Classic, use the ACL ID, for example `aws_waf_web_acl.example.id`. The WAF Web ACL must exist in the WAF Global (CloudFront) region and the credentials configuring this argument must have `waf:GetWebACL` permissions assigned.
	WebAclId string `json:"webAclId,omitempty" yaml:"webAclId,omitempty"`

	// Any comments you want to include about the distribution.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// Price class for this distribution. One of `PriceClass_All`, `PriceClass_200`, `PriceClass_100`.
	PriceClass string `json:"priceClass,omitempty" yaml:"priceClass,omitempty"`

	// A Boolean that indicates whether this is a staging distribution. Defaults to `false`.
	Staging bool `json:"staging,omitempty" yaml:"staging,omitempty"`

	// Ordered list of cache behaviors resource for this distribution. List from top to bottom in order of precedence. The topmost cache behavior will have precedence 0.
	OrderedCacheBehaviors []types.Cloudfront_DistributionOrderedCacheBehavior `json:"orderedCacheBehaviors,omitempty" yaml:"orderedCacheBehaviors,omitempty"`

	// One or more origin_group for this distribution (multiples allowed).
	OriginGroups []types.Cloudfront_DistributionOriginGroup `json:"originGroups,omitempty" yaml:"originGroups,omitempty"`

	// Disables the distribution instead of deleting it when destroying the resource through the provider. If this is set, the distribution needs to be deleted manually afterwards. Default: `false`.
	RetainOnDelete bool `json:"retainOnDelete,omitempty" yaml:"retainOnDelete,omitempty"`

	// Object that you want CloudFront to return (for example, index.html) when an end user requests the root URL.
	DefaultRootObject string `json:"defaultRootObject,omitempty" yaml:"defaultRootObject,omitempty"`

	// Whether Origin Shield is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Whether the IPv6 is enabled for the distribution.
	IsIpv6Enabled bool `json:"isIpv6Enabled,omitempty" yaml:"isIpv6Enabled,omitempty"`

	// Extra CNAMEs (alternate domain names), if any, for this distribution.
	Aliases []string `json:"aliases,omitempty" yaml:"aliases,omitempty"`

	// One or more custom error response elements (multiples allowed).
	CustomErrorResponses []types.Cloudfront_DistributionCustomErrorResponse `json:"customErrorResponses,omitempty" yaml:"customErrorResponses,omitempty"`

	// Maximum HTTP version to support on the distribution. Allowed values are `http1.1`, `http2`, `http2and3` and `http3`. The default is `http2`.
	HttpVersion string `json:"httpVersion,omitempty" yaml:"httpVersion,omitempty"`
}
