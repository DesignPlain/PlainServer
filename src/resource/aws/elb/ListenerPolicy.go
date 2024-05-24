package elb

type ListenerPolicy struct {
	// The load balancer to attach the policy to.
	LoadBalancerName string `json:"loadBalancerName,omitempty" yaml:"loadBalancerName,omitempty"`

	// The load balancer listener port to apply the policy to.
	LoadBalancerPort int `json:"loadBalancerPort,omitempty" yaml:"loadBalancerPort,omitempty"`

	// List of Policy Names to apply to the backend server.
	PolicyNames []string `json:"policyNames,omitempty" yaml:"policyNames,omitempty"`

	// Map of arbitrary keys and values that, when changed, will trigger an update.
	Triggers map[string]string `json:"triggers,omitempty" yaml:"triggers,omitempty"`
}
