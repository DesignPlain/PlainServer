package types

type Route53_RecordFailoverRoutingPolicy struct {
	// `PRIMARY` or `SECONDARY`. A `PRIMARY` record will be served if its healthcheck is passing, otherwise the `SECONDARY` will be served. See http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-configuring-options.html#dns-failover-failover-rrsets
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
