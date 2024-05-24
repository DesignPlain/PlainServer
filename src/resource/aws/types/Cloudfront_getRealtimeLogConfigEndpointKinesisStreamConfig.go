package types

type Cloudfront_getRealtimeLogConfigEndpointKinesisStreamConfig struct {
	/*
	   (Required) ARN of an IAM role that CloudFront can use to send real-time log data to the Kinesis data stream.
	   See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-iam-role) for more information.
	*/
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// (Required) ARN of the Kinesis data stream.
	StreamArn string `json:"streamArn,omitempty" yaml:"streamArn,omitempty"`
}
