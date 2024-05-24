package rum

type MetricsDestination struct {
	// Defines the destination to send the metrics to. Valid values are `CloudWatch` and `Evidently`. If you specify `Evidently`, you must also specify the ARN of the CloudWatchEvidently experiment that is to be the destination and an IAM role that has permission to write to the experiment.
	Destination string `json:"destination,omitempty" yaml:"destination,omitempty"`

	// Use this parameter only if Destination is Evidently. This parameter specifies the ARN of the Evidently experiment that will receive the extended metrics.
	DestinationArn string `json:"destinationArn,omitempty" yaml:"destinationArn,omitempty"`

	// This parameter is required if Destination is Evidently. If Destination is CloudWatch, do not use this parameter.
	IamRoleArn string `json:"iamRoleArn,omitempty" yaml:"iamRoleArn,omitempty"`

	// The name of the CloudWatch RUM app monitor that will send the metrics.
	AppMonitorName string `json:"appMonitorName,omitempty" yaml:"appMonitorName,omitempty"`
}
