package sns

type SmsPreferences struct {
	// The ARN of the IAM role that allows Amazon SNS to write logs about SMS deliveries in CloudWatch Logs.
	DeliveryStatusIamRoleArn string `json:"deliveryStatusIamRoleArn,omitempty" yaml:"deliveryStatusIamRoleArn,omitempty"`

	// The percentage of successful SMS deliveries for which Amazon SNS will write logs in CloudWatch Logs. The value must be between 0 and 100.
	DeliveryStatusSuccessSamplingRate string `json:"deliveryStatusSuccessSamplingRate,omitempty" yaml:"deliveryStatusSuccessSamplingRate,omitempty"`

	// The maximum amount in USD that you are willing to spend each month to send SMS messages.
	MonthlySpendLimit int `json:"monthlySpendLimit,omitempty" yaml:"monthlySpendLimit,omitempty"`

	// The name of the Amazon S3 bucket to receive daily SMS usage reports from Amazon SNS.
	UsageReportS3Bucket string `json:"usageReportS3Bucket,omitempty" yaml:"usageReportS3Bucket,omitempty"`

	// A string, such as your business brand, that is displayed as the sender on the receiving device.
	DefaultSenderId string `json:"defaultSenderId,omitempty" yaml:"defaultSenderId,omitempty"`

	// The type of SMS message that you will send by default. Possible values are: Promotional, Transactional
	DefaultSmsType string `json:"defaultSmsType,omitempty" yaml:"defaultSmsType,omitempty"`
}
