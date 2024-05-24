package types

type Ecs_TaskSetLoadBalancer struct {
	// The name of the container to associate with the load balancer (as it appears in a container definition).
	ContainerName string `json:"containerName,omitempty" yaml:"containerName,omitempty"`

	/*
	   The port on the container to associate with the load balancer. Defaults to `0` if not specified.

	   > --Note:-- Specifying multiple `load_balancer` configurations is still not supported by AWS for ECS task set.
	*/
	ContainerPort int `json:"containerPort,omitempty" yaml:"containerPort,omitempty"`

	// The name of the ELB (Classic) to associate with the service.
	LoadBalancerName string `json:"loadBalancerName,omitempty" yaml:"loadBalancerName,omitempty"`

	// The ARN of the Load Balancer target group to associate with the service.
	TargetGroupArn string `json:"targetGroupArn,omitempty" yaml:"targetGroupArn,omitempty"`
}
