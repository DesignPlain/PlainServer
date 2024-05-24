package types

type Lb_ListenerDefaultActionForwardStickiness struct {
	/*
	   Time period, in seconds, during which requests from a client should be routed to the same target group. The range is 1-604800 seconds (7 days).

	   The following arguments are optional:
	*/
	Duration int `json:"duration,omitempty" yaml:"duration,omitempty"`

	// Whether target group stickiness is enabled. Default is `false`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
