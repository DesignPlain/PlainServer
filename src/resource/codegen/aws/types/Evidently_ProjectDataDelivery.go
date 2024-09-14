package types

type Evidently_ProjectDataDelivery struct {
	// A block that defines the CloudWatch Log Group that stores the evaluation events. See below.
	CloudwatchLogs Evidently_ProjectDataDeliveryCloudwatchLogs `json:"cloudwatchLogs,omitempty" yaml:"cloudwatchLogs,omitempty"`

	// A block that defines the S3 bucket and prefix that stores the evaluation events. See below.
	S3Destination Evidently_ProjectDataDeliveryS3Destination `json:"s3Destination,omitempty" yaml:"s3Destination,omitempty"`
}
