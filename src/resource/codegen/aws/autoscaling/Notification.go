package autoscaling

type Notification struct {
	// List of AutoScaling Group Names
	GroupNames []string `json:"groupNames,omitempty" yaml:"groupNames,omitempty"`

	/*
	   List of Notification Types that trigger
	   notifications. Acceptable values are documented [in the AWS documentation here](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_NotificationConfiguration.html)
	*/
	Notifications []string `json:"notifications,omitempty" yaml:"notifications,omitempty"`

	// Topic ARN for notifications to be sent through
	TopicArn string `json:"topicArn,omitempty" yaml:"topicArn,omitempty"`
}
