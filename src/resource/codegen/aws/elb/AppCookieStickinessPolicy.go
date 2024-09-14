package elb

type AppCookieStickinessPolicy struct {
	// Name of the stickiness policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Application cookie whose lifetime the ELB's cookie should follow.
	CookieName string `json:"cookieName,omitempty" yaml:"cookieName,omitempty"`

	/*
	   Load balancer port to which the policy
	   should be applied. This must be an active listener on the load
	   balancer.
	*/
	LbPort int `json:"lbPort,omitempty" yaml:"lbPort,omitempty"`

	/*
	   Name of load balancer to which the policy
	   should be attached.
	*/
	LoadBalancer string `json:"loadBalancer,omitempty" yaml:"loadBalancer,omitempty"`
}
