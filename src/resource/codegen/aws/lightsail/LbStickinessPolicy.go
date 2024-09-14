package lightsail

type LbStickinessPolicy struct {
	// The cookie duration in seconds. This determines the length of the session stickiness.
	CookieDuration int `json:"cookieDuration,omitempty" yaml:"cookieDuration,omitempty"`

	// The Session Stickiness state of the load balancer. `true` to activate session stickiness or `false` to deactivate session stickiness.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The name of the load balancer to which you want to enable session stickiness.
	LbName string `json:"lbName,omitempty" yaml:"lbName,omitempty"`
}
