package types

type Cloudwatch_EventEndpointRoutingConfig struct {
	// Parameters used for failover. This includes what triggers failover and what happens when it's triggered. Documented below.
	FailoverConfig Cloudwatch_EventEndpointRoutingConfigFailoverConfig `json:"failoverConfig,omitempty" yaml:"failoverConfig,omitempty"`
}
