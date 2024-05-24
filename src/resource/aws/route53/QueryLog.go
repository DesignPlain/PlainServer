package route53

type QueryLog struct {
	// CloudWatch log group ARN to send query logs.
	CloudwatchLogGroupArn string `json:"cloudwatchLogGroupArn,omitempty" yaml:"cloudwatchLogGroupArn,omitempty"`

	// Route53 hosted zone ID to enable query logs.
	ZoneId string `json:"zoneId,omitempty" yaml:"zoneId,omitempty"`
}
