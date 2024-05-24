package elb

type LoadBalancerBackendServerPolicy struct {
	// The load balancer to attach the policy to.
	LoadBalancerName string `json:"loadBalancerName,omitempty" yaml:"loadBalancerName,omitempty"`

	// List of Policy Names to apply to the backend server.
	PolicyNames []string `json:"policyNames,omitempty" yaml:"policyNames,omitempty"`

	// The instance port to apply the policy to.
	InstancePort int `json:"instancePort,omitempty" yaml:"instancePort,omitempty"`
}
