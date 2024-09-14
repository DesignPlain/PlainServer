package types

type Cloudwatch_EventEndpointRoutingConfigFailoverConfig struct {
	// Parameters used for the primary Region. Documented below.
	Primary Cloudwatch_EventEndpointRoutingConfigFailoverConfigPrimary `json:"primary,omitempty" yaml:"primary,omitempty"`

	// Parameters used for the secondary Region, the Region that events are routed to when failover is triggered or event replication is enabled. Documented below.
	Secondary Cloudwatch_EventEndpointRoutingConfigFailoverConfigSecondary `json:"secondary,omitempty" yaml:"secondary,omitempty"`
}
