package s3

import types "libds/aws/types"

type BucketWebsiteConfigurationV2 struct {
	// List of rules that define when a redirect is applied and the redirect behavior. See below.
	RoutingRules []types.S3_BucketWebsiteConfigurationV2RoutingRule `json:"routingRules,omitempty" yaml:"routingRules,omitempty"`

	// Name of the bucket.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Name of the error document for the website. See below.
	ErrorDocument types.S3_BucketWebsiteConfigurationV2ErrorDocument `json:"errorDocument,omitempty" yaml:"errorDocument,omitempty"`

	// Account ID of the expected bucket owner.
	ExpectedBucketOwner string `json:"expectedBucketOwner,omitempty" yaml:"expectedBucketOwner,omitempty"`

	// Name of the index document for the website. See below.
	IndexDocument types.S3_BucketWebsiteConfigurationV2IndexDocument `json:"indexDocument,omitempty" yaml:"indexDocument,omitempty"`

	// Redirect behavior for every request to this bucket's website endpoint. See below. Conflicts with `error_document`, `index_document`, and `routing_rule`.
	RedirectAllRequestsTo types.S3_BucketWebsiteConfigurationV2RedirectAllRequestsTo `json:"redirectAllRequestsTo,omitempty" yaml:"redirectAllRequestsTo,omitempty"`

	/*
	   JSON array containing [routing rules](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html)
	   describing redirect behavior and when redirects are applied. Use this parameter when your routing rules contain empty String values (`""`) as seen in the example above.
	*/
	RoutingRuleDetails string `json:"routingRuleDetails,omitempty" yaml:"routingRuleDetails,omitempty"`
}
