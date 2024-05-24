package types

type Cloudfront_DistributionOrderedCacheBehaviorFunctionAssociation struct {
	// Specific event to trigger this function. Valid values: `viewer-request` or `viewer-response`.
	EventType string `json:"eventType,omitempty" yaml:"eventType,omitempty"`

	// ARN of the CloudFront function.
	FunctionArn string `json:"functionArn,omitempty" yaml:"functionArn,omitempty"`
}
