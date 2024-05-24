package sns

type PlatformApplication struct {
	// The identifier that's assigned to your Apple developer account team. Must be 10 alphanumeric characters.
	ApplePlatformTeamId string `json:"applePlatformTeamId,omitempty" yaml:"applePlatformTeamId,omitempty"`

	// The ARN of the SNS Topic triggered when a delivery to any of the platform endpoints associated with your platform application encounters a permanent failure.
	EventDeliveryFailureTopicArn string `json:"eventDeliveryFailureTopicArn,omitempty" yaml:"eventDeliveryFailureTopicArn,omitempty"`

	// The ARN of the SNS Topic triggered when a new platform endpoint is added to your platform application.
	EventEndpointCreatedTopicArn string `json:"eventEndpointCreatedTopicArn,omitempty" yaml:"eventEndpointCreatedTopicArn,omitempty"`

	// The ARN of the SNS Topic triggered when an existing platform endpoint is deleted from your platform application.
	EventEndpointDeletedTopicArn string `json:"eventEndpointDeletedTopicArn,omitempty" yaml:"eventEndpointDeletedTopicArn,omitempty"`

	// The friendly name for the SNS platform application
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The platform that the app is registered with. See [Platform](http://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-register.html) for supported platforms.
	Platform string `json:"platform,omitempty" yaml:"platform,omitempty"`

	// Application Platform credential. See [Credential](http://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-register.html) for type of credential required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
	PlatformCredential string `json:"platformCredential,omitempty" yaml:"platformCredential,omitempty"`

	// The bundle identifier that's assigned to your iOS app. May only include alphanumeric characters, hyphens (-), and periods (.).
	ApplePlatformBundleId string `json:"applePlatformBundleId,omitempty" yaml:"applePlatformBundleId,omitempty"`

	// The IAM role ARN permitted to receive success feedback for this application and give SNS write access to use CloudWatch logs on your behalf.
	SuccessFeedbackRoleArn string `json:"successFeedbackRoleArn,omitempty" yaml:"successFeedbackRoleArn,omitempty"`

	// Application Platform principal. See [Principal](http://docs.aws.amazon.com/sns/latest/api/API_CreatePlatformApplication.html) for type of principal required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
	PlatformPrincipal string `json:"platformPrincipal,omitempty" yaml:"platformPrincipal,omitempty"`

	// The IAM role ARN permitted to receive failure feedback for this application and give SNS write access to use CloudWatch logs on your behalf.
	FailureFeedbackRoleArn string `json:"failureFeedbackRoleArn,omitempty" yaml:"failureFeedbackRoleArn,omitempty"`

	/*
	   The sample rate percentage (0-100) of successfully delivered messages.

	   The following attributes are needed only when using APNS token credentials:
	*/
	SuccessFeedbackSampleRate string `json:"successFeedbackSampleRate,omitempty" yaml:"successFeedbackSampleRate,omitempty"`

	// The ARN of the SNS Topic triggered when an existing platform endpoint is changed from your platform application.
	EventEndpointUpdatedTopicArn string `json:"eventEndpointUpdatedTopicArn,omitempty" yaml:"eventEndpointUpdatedTopicArn,omitempty"`
}
