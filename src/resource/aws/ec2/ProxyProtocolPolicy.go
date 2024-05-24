package ec2

type ProxyProtocolPolicy struct {
	/*
	   The load balancer to which the policy
	   should be attached.
	*/
	LoadBalancer string `json:"loadBalancer,omitempty" yaml:"loadBalancer,omitempty"`

	/*
	   List of instance ports to which the policy
	   should be applied. This can be specified if the protocol is SSL or TCP.
	*/
	InstancePorts []string `json:"instancePorts,omitempty" yaml:"instancePorts,omitempty"`
}
