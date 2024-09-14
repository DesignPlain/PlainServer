package types

type Route53_RecordWeightedRoutingPolicy struct {
	// A numeric value indicating the relative weight of the record. See http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html#routing-policy-weighted.
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`
}
