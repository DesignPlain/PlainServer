package types

type Sagemaker_EndpointConfigurationAsyncInferenceConfigOutputConfigNotificationConfig struct {
	// The Amazon SNS topics where you want the inference response to be included. Valid values are `SUCCESS_NOTIFICATION_TOPIC` and `ERROR_NOTIFICATION_TOPIC`.
	IncludeInferenceResponseIns []string `json:"includeInferenceResponseIns,omitempty" yaml:"includeInferenceResponseIns,omitempty"`

	// Amazon SNS topic to post a notification to when inference completes successfully. If no topic is provided, no notification is sent on success.
	SuccessTopic string `json:"successTopic,omitempty" yaml:"successTopic,omitempty"`

	// Amazon SNS topic to post a notification to when inference fails. If no topic is provided, no notification is sent on failure.
	ErrorTopic string `json:"errorTopic,omitempty" yaml:"errorTopic,omitempty"`
}
