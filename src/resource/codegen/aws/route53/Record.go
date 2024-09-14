package route53

import types "libds/aws/types"

type Record struct {
	// A block indicating a routing policy based on the geoproximity of the requestor. Conflicts with any other routing policy. Documented below.
	GeoproximityRoutingPolicy types.Route53_RecordGeoproximityRoutingPolicy `json:"geoproximityRoutingPolicy,omitempty" yaml:"geoproximityRoutingPolicy,omitempty"`

	// The name of the record.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Unique identifier to differentiate records with routing policies from one another. Required if using `cidr_routing_policy`, `failover_routing_policy`, `geolocation_routing_policy`,`geoproximity_routing_policy`, `latency_routing_policy`, `multivalue_answer_routing_policy`, or `weighted_routing_policy`.
	SetIdentifier string `json:"setIdentifier,omitempty" yaml:"setIdentifier,omitempty"`

	// The record type. Valid values are `A`, `AAAA`, `CAA`, `CNAME`, `DS`, `MX`, `NAPTR`, `NS`, `PTR`, `SOA`, `SPF`, `SRV` and `TXT`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The health check the record should be associated with.
	HealthCheckId string `json:"healthCheckId,omitempty" yaml:"healthCheckId,omitempty"`

	// A string list of records. To specify a single record value longer than 255 characters such as a TXT record for DKIM, add `\"\"` inside the provider configuration string (e.g., `"first255characters\"\"morecharacters"`).
	Records []string `json:"records,omitempty" yaml:"records,omitempty"`

	// The ID of the hosted zone to contain this record.
	ZoneId string `json:"zoneId,omitempty" yaml:"zoneId,omitempty"`

	// A block indicating the routing behavior when associated health check fails. Conflicts with any other routing policy. Documented below.
	FailoverRoutingPolicies []types.Route53_RecordFailoverRoutingPolicy `json:"failoverRoutingPolicies,omitempty" yaml:"failoverRoutingPolicies,omitempty"`

	// The TTL of the record.
	Ttl int `json:"ttl,omitempty" yaml:"ttl,omitempty"`

	/*
	   An alias block. Conflicts with `ttl` & `records`.
	   Documented below.
	*/
	Aliases []types.Route53_RecordAlias `json:"aliases,omitempty" yaml:"aliases,omitempty"`

	/*
	   Allow creation of this record to overwrite an existing record, if any. This does not affect the ability to update the record using this provider and does not prevent other resources within this provider or manual Route 53 changes outside this provider from overwriting this record. `false` by default. This configuration is not recommended for most environments.

	   Exactly one of `records` or `alias` must be specified: this determines whether it's an alias record.
	*/
	AllowOverwrite bool `json:"allowOverwrite,omitempty" yaml:"allowOverwrite,omitempty"`

	// A block indicating a routing policy based on the IP network ranges of requestors. Conflicts with any other routing policy. Documented below.
	CidrRoutingPolicy types.Route53_RecordCidrRoutingPolicy `json:"cidrRoutingPolicy,omitempty" yaml:"cidrRoutingPolicy,omitempty"`

	// A block indicating a routing policy based on the geolocation of the requestor. Conflicts with any other routing policy. Documented below.
	GeolocationRoutingPolicies []types.Route53_RecordGeolocationRoutingPolicy `json:"geolocationRoutingPolicies,omitempty" yaml:"geolocationRoutingPolicies,omitempty"`

	// A block indicating a routing policy based on the latency between the requestor and an AWS region. Conflicts with any other routing policy. Documented below.
	LatencyRoutingPolicies []types.Route53_RecordLatencyRoutingPolicy `json:"latencyRoutingPolicies,omitempty" yaml:"latencyRoutingPolicies,omitempty"`

	// Set to `true` to indicate a multivalue answer routing policy. Conflicts with any other routing policy.
	MultivalueAnswerRoutingPolicy bool `json:"multivalueAnswerRoutingPolicy,omitempty" yaml:"multivalueAnswerRoutingPolicy,omitempty"`

	// A block indicating a weighted routing policy. Conflicts with any other routing policy. Documented below.
	WeightedRoutingPolicies []types.Route53_RecordWeightedRoutingPolicy `json:"weightedRoutingPolicies,omitempty" yaml:"weightedRoutingPolicies,omitempty"`
}
