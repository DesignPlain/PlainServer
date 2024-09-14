package types

type Lb_ListenerRuleActionForwardStickiness struct {
	// The time period, in seconds, during which requests from a client should be routed to the same target group. The range is 1-604800 seconds (7 days).
	Duration int `json:"duration,omitempty" yaml:"duration,omitempty"`

	// Indicates whether target group stickiness is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
