package ses

type IdentityNotificationTopic struct {
	// The identity for which the Amazon SNS topic will be set. You can specify an identity by using its name or by using its Amazon Resource Name (ARN).
	Identity string `json:"identity,omitempty" yaml:"identity,omitempty"`

	// Whether SES should include original email headers in SNS notifications of this type. `false` by default.
	IncludeOriginalHeaders bool `json:"includeOriginalHeaders,omitempty" yaml:"includeOriginalHeaders,omitempty"`

	// The type of notifications that will be published to the specified Amazon SNS topic. Valid Values: `Bounce`, `Complaint` or `Delivery`.
	NotificationType string `json:"notificationType,omitempty" yaml:"notificationType,omitempty"`

	// The Amazon Resource Name (ARN) of the Amazon SNS topic. Can be set to `""` (an empty string) to disable publishing.
	TopicArn string `json:"topicArn,omitempty" yaml:"topicArn,omitempty"`
}
