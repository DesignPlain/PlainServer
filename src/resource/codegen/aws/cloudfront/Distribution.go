package cloudfront

import types "libds/aws/types"

type Distribution struct {
	//
	DefaultCacheBehavior types.Cloudfront_DistributionDefaultCacheBehavior `json:"defaultCacheBehavior,omitempty" yaml:"defaultCacheBehavior,omitempty"`

	//
	OrderedCacheBehaviors []types.Cloudfront_DistributionOrderedCacheBehavior `json:"orderedCacheBehaviors,omitempty" yaml:"orderedCacheBehaviors,omitempty"`

	//
	OriginGroups []types.Cloudfront_DistributionOriginGroup `json:"originGroups,omitempty" yaml:"originGroups,omitempty"`

	//
	WaitForDeployment bool `json:"waitForDeployment,omitempty" yaml:"waitForDeployment,omitempty"`

	//
	WebAclId string `json:"webAclId,omitempty" yaml:"webAclId,omitempty"`

	//
	CustomErrorResponses []types.Cloudfront_DistributionCustomErrorResponse `json:"customErrorResponses,omitempty" yaml:"customErrorResponses,omitempty"`

	//
	Origins []types.Cloudfront_DistributionOrigin `json:"origins,omitempty" yaml:"origins,omitempty"`

	//
	ContinuousDeploymentPolicyId string `json:"continuousDeploymentPolicyId,omitempty" yaml:"continuousDeploymentPolicyId,omitempty"`

	//
	DefaultRootObject string `json:"defaultRootObject,omitempty" yaml:"defaultRootObject,omitempty"`

	//
	IsIpv6Enabled bool `json:"isIpv6Enabled,omitempty" yaml:"isIpv6Enabled,omitempty"`

	//
	PriceClass string `json:"priceClass,omitempty" yaml:"priceClass,omitempty"`

	//
	RetainOnDelete bool `json:"retainOnDelete,omitempty" yaml:"retainOnDelete,omitempty"`

	//
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// `true` if any of the AWS accounts listed as trusted signers have active CloudFront key pairs
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	//
	HttpVersion string `json:"httpVersion,omitempty" yaml:"httpVersion,omitempty"`

	//
	LoggingConfig types.Cloudfront_DistributionLoggingConfig `json:"loggingConfig,omitempty" yaml:"loggingConfig,omitempty"`

	//
	Restrictions types.Cloudfront_DistributionRestrictions `json:"restrictions,omitempty" yaml:"restrictions,omitempty"`

	//
	Staging bool `json:"staging,omitempty" yaml:"staging,omitempty"`

	//
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	ViewerCertificate types.Cloudfront_DistributionViewerCertificate `json:"viewerCertificate,omitempty" yaml:"viewerCertificate,omitempty"`

	//
	Aliases []string `json:"aliases,omitempty" yaml:"aliases,omitempty"`
}
