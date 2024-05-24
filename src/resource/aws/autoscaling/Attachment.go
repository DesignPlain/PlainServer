package autoscaling

type Attachment struct {
	// Name of ASG to associate with the ELB.
	AutoscalingGroupName string `json:"autoscalingGroupName,omitempty" yaml:"autoscalingGroupName,omitempty"`

	// Name of the ELB.
	Elb string `json:"elb,omitempty" yaml:"elb,omitempty"`

	// ARN of a load balancer target group.
	LbTargetGroupArn string `json:"lbTargetGroupArn,omitempty" yaml:"lbTargetGroupArn,omitempty"`
}
