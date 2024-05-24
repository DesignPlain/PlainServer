package types

type Route53_getTrafficPolicyDocumentEndpoint struct {
	// ID of an endpoint you want to assign.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// To route traffic to an Amazon S3 bucket that is configured as a website endpoint, specify the region in which you created the bucket for `region`.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Type of the endpoint. Valid values are `value` , `cloudfront` , `elastic-load-balancer`, `s3-website`
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Value of the `type`.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
