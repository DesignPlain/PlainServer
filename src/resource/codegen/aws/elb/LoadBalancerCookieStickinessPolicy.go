package elb

type LoadBalancerCookieStickinessPolicy struct {
	/*
	   The time period after which
	   the session cookie should be considered stale, expressed in seconds.
	*/
	CookieExpirationPeriod int `json:"cookieExpirationPeriod,omitempty" yaml:"cookieExpirationPeriod,omitempty"`

	/*
	   The load balancer port to which the policy
	   should be applied. This must be an active listener on the load
	   balancer.
	*/
	LbPort int `json:"lbPort,omitempty" yaml:"lbPort,omitempty"`

	/*
	   The load balancer to which the policy
	   should be attached.
	*/
	LoadBalancer string `json:"loadBalancer,omitempty" yaml:"loadBalancer,omitempty"`

	// The name of the stickiness policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
