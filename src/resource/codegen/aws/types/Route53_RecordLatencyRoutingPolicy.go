package types

type Route53_RecordLatencyRoutingPolicy struct {
	// An AWS region from which to measure latency. See http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html#routing-policy-latency
	Region string `json:"region,omitempty" yaml:"region,omitempty"`
}
