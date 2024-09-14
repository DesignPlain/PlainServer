package types

type Alb_TargetGroupStickiness struct {
	// Name of the application based cookie. AWSALB, AWSALBAPP, and AWSALBTG prefixes are reserved and cannot be used. Only needed when type is `app_cookie`.
	CookieName string `json:"cookieName,omitempty" yaml:"cookieName,omitempty"`

	// Boolean to enable / disable `stickiness`. Default is `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The type of sticky sessions. The only current possible values are `lb_cookie`, `app_cookie` for ALBs, `source_ip` for NLBs, and `source_ip_dest_ip`, `source_ip_dest_ip_proto` for GWLBs.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Only used when the type is `lb_cookie`. The time period, in seconds, during which requests from a client should be routed to the same target. After this time period expires, the load balancer-generated cookie is considered stale. The range is 1 second to 1 week (604800 seconds). The default value is 1 day (86400 seconds).
	CookieDuration int `json:"cookieDuration,omitempty" yaml:"cookieDuration,omitempty"`
}
