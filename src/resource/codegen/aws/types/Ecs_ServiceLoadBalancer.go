package types

type Ecs_ServiceLoadBalancer struct {
	// Name of the ELB (Classic) to associate with the service.
	ElbName string `json:"elbName,omitempty" yaml:"elbName,omitempty"`

	// ARN of the Load Balancer target group to associate with the service.
	TargetGroupArn string `json:"targetGroupArn,omitempty" yaml:"targetGroupArn,omitempty"`

	// Name of the container to associate with the load balancer (as it appears in a container definition).
	ContainerName string `json:"containerName,omitempty" yaml:"containerName,omitempty"`

	/*
	   Port on the container to associate with the load balancer.

	   > --Version note:-- Multiple `load_balancer` configuration block support was added in version 2.22.0 of the provider. This allows configuration of [ECS service support for multiple target groups](https://aws.amazon.com/about-aws/whats-new/2019/07/amazon-ecs-services-now-support-multiple-load-balancer-target-groups/).
	*/
	ContainerPort int `json:"containerPort,omitempty" yaml:"containerPort,omitempty"`
}
