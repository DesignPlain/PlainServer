package types

type Route53_RecordAlias struct {
	// Set to `true` if you want Route 53 to determine whether to respond to DNS queries using this resource record set by checking the health of the resource record set. Some resources have special requirements, see [related part of documentation](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resource-record-sets-values.html#rrsets-values-alias-evaluate-target-health).
	EvaluateTargetHealth bool `json:"evaluateTargetHealth,omitempty" yaml:"evaluateTargetHealth,omitempty"`

	// DNS domain name for a CloudFront distribution, S3 bucket, ELB, or another resource record set in this hosted zone.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Hosted zone ID for a CloudFront distribution, S3 bucket, ELB, or Route 53 hosted zone. See `resource_elb.zone_id` for example.
	ZoneId string `json:"zoneId,omitempty" yaml:"zoneId,omitempty"`
}
