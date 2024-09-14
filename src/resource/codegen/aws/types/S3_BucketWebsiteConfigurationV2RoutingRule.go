package types

type S3_BucketWebsiteConfigurationV2RoutingRule struct {
	// Configuration block for describing a condition that must be met for the specified redirect to apply. See below.
	Condition S3_BucketWebsiteConfigurationV2RoutingRuleCondition `json:"condition,omitempty" yaml:"condition,omitempty"`

	// Configuration block for redirect information. See below.
	Redirect S3_BucketWebsiteConfigurationV2RoutingRuleRedirect `json:"redirect,omitempty" yaml:"redirect,omitempty"`
}
