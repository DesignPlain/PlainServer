package types

type S3_BucketNotificationTopic struct {
	// [Event](http://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html#notification-how-to-event-types-and-destinations) for which to send notifications.
	Events []string `json:"events,omitempty" yaml:"events,omitempty"`

	// Object key name prefix.
	FilterPrefix string `json:"filterPrefix,omitempty" yaml:"filterPrefix,omitempty"`

	// Object key name suffix.
	FilterSuffix string `json:"filterSuffix,omitempty" yaml:"filterSuffix,omitempty"`

	// Unique identifier for each of the notification configurations.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// SNS topic ARN.
	TopicArn string `json:"topicArn,omitempty" yaml:"topicArn,omitempty"`
}
