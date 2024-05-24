package types

type Vpclattice_TargetGroupAttachmentTarget struct {
	// The ID of the target. If the target type of the target group is INSTANCE, this is an instance ID. If the target type is IP , this is an IP address. If the target type is LAMBDA, this is the ARN of the Lambda function. If the target type is ALB, this is the ARN of the Application Load Balancer.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// This port is used for routing traffic to the target, and defaults to the target group port. However, you can override the default and specify a custom port.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
