package types

type Securitylake_SubscriberNotificationConfiguration struct {
	// The configurations for HTTPS subscriber notification.
	HttpsNotificationConfiguration Securitylake_SubscriberNotificationConfigurationHttpsNotificationConfiguration `json:"httpsNotificationConfiguration,omitempty" yaml:"httpsNotificationConfiguration,omitempty"`

	/*
	   The configurations for SQS subscriber notification.
	   There are no parameters within `sqs_notification_configuration`.
	*/
	SqsNotificationConfiguration Securitylake_SubscriberNotificationConfigurationSqsNotificationConfiguration `json:"sqsNotificationConfiguration,omitempty" yaml:"sqsNotificationConfiguration,omitempty"`
}
