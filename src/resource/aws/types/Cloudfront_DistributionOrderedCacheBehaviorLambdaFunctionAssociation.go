package types

type Cloudfront_DistributionOrderedCacheBehaviorLambdaFunctionAssociation struct {
	// When set to true it exposes the request body to the lambda function. Defaults to false. Valid values: `true`, `false`.
	IncludeBody bool `json:"includeBody,omitempty" yaml:"includeBody,omitempty"`

	// ARN of the Lambda function.
	LambdaArn string `json:"lambdaArn,omitempty" yaml:"lambdaArn,omitempty"`

	// Specific event to trigger this function. Valid values: `viewer-request`, `origin-request`, `viewer-response`, `origin-response`.
	EventType string `json:"eventType,omitempty" yaml:"eventType,omitempty"`
}
