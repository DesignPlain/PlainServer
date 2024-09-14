package types

type Rekognition_StreamProcessorNotificationChannel struct {
	// The Amazon Resource Number (ARN) of the Amazon Amazon Simple Notification Service topic to which Amazon Rekognition posts the completion status.
	SnsTopicArn string `json:"snsTopicArn,omitempty" yaml:"snsTopicArn,omitempty"`
}
