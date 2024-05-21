package types

type Container_ClusterAddonsConfigCloudrunConfig struct {
	// The status of the CloudRun addon. It is disabled by default. Set `disabled=false` to enable.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	/*
	   The load balancer type of CloudRun ingress service. It is external load balancer by default.
	   Set `load_balancer_type=LOAD_BALANCER_TYPE_INTERNAL` to configure it as internal load balancer.
	*/
	LoadBalancerType string `json:"loadBalancerType,omitempty" yaml:"loadBalancerType,omitempty"`
}
