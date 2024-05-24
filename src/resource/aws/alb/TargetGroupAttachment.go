package alb

type TargetGroupAttachment struct {
	// The port on which targets receive traffic.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// The ARN of the target group with which to register targets.
	TargetGroupArn string `json:"targetGroupArn,omitempty" yaml:"targetGroupArn,omitempty"`

	/*
	   The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is `ip`, specify an IP address. If the target type is `lambda`, specify the Lambda function ARN. If the target type is `alb`, specify the ALB ARN.

	   The following arguments are optional:
	*/
	TargetId string `json:"targetId,omitempty" yaml:"targetId,omitempty"`

	// The Availability Zone where the IP address of the target is to be registered. If the private IP address is outside of the VPC scope, this value must be set to `all`.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`
}
