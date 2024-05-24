package types

type Autoscaling_PolicyPredictiveScalingConfigurationMetricSpecificationPredefinedLoadMetricSpecification struct {
	// Metric type. Valid values are `ASGTotalCPUUtilization`, `ASGTotalNetworkIn`, `ASGTotalNetworkOut`, or `ALBTargetGroupRequestCount`.
	PredefinedMetricType string `json:"predefinedMetricType,omitempty" yaml:"predefinedMetricType,omitempty"`

	// Label that uniquely identifies a specific Application Load Balancer target group from which to determine the request count served by your Auto Scaling group. You create the resource label by appending the final portion of the load balancer ARN and the final portion of the target group ARN into a single value, separated by a forward slash (/). Refer to [PredefinedMetricSpecification](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_PredefinedMetricSpecification.html) for more information.
	ResourceLabel string `json:"resourceLabel,omitempty" yaml:"resourceLabel,omitempty"`
}
