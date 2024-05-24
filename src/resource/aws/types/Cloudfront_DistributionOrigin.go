package types

type Cloudfront_DistributionOrigin struct {
	// Optional element that causes CloudFront to request your content from a directory in your Amazon S3 bucket or your custom origin.
	OriginPath string `json:"originPath,omitempty" yaml:"originPath,omitempty"`

	// CloudFront S3 origin configuration information. If a custom origin is required, use `custom_origin_config` instead.
	S3OriginConfig Cloudfront_DistributionOriginS3OriginConfig `json:"s3OriginConfig,omitempty" yaml:"s3OriginConfig,omitempty"`

	// Number of seconds that CloudFront waits when trying to establish a connection to the origin. Must be between 1-10. Defaults to 10.
	ConnectionTimeout int `json:"connectionTimeout,omitempty" yaml:"connectionTimeout,omitempty"`

	// One or more sub-resources with `name` and `value` parameters that specify header data that will be sent to the origin (multiples allowed).
	CustomHeaders []Cloudfront_DistributionOriginCustomHeader `json:"customHeaders,omitempty" yaml:"customHeaders,omitempty"`

	// DNS domain name of either the S3 bucket, or web site of your custom origin.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// Unique identifier of a [CloudFront origin access control][8] for this origin.
	OriginAccessControlId string `json:"originAccessControlId,omitempty" yaml:"originAccessControlId,omitempty"`

	// Unique identifier of the member origin.
	OriginId string `json:"originId,omitempty" yaml:"originId,omitempty"`

	// CloudFront Origin Shield configuration information. Using Origin Shield can help reduce the load on your origin. For more information, see [Using Origin Shield](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html) in the Amazon CloudFront Developer Guide.
	OriginShield Cloudfront_DistributionOriginOriginShield `json:"originShield,omitempty" yaml:"originShield,omitempty"`

	// Number of times that CloudFront attempts to connect to the origin. Must be between 1-3. Defaults to 3.
	ConnectionAttempts int `json:"connectionAttempts,omitempty" yaml:"connectionAttempts,omitempty"`

	// The CloudFront custom origin configuration information. If an S3 origin is required, use `origin_access_control_id` or `s3_origin_config` instead.
	CustomOriginConfig Cloudfront_DistributionOriginCustomOriginConfig `json:"customOriginConfig,omitempty" yaml:"customOriginConfig,omitempty"`
}
