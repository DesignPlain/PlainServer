package s3

import types "DesignSphere_Server/src/resource/aws/types"

type BucketNotification struct {
	/*
	   Name of the bucket for notification configuration.

	   The following arguments are optional:
	*/
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Whether to enable Amazon EventBridge notifications. Defaults to `false`.
	Eventbridge bool `json:"eventbridge,omitempty" yaml:"eventbridge,omitempty"`

	// Used to configure notifications to a Lambda Function. See below.
	LambdaFunctions []types.S3_BucketNotificationLambdaFunction `json:"lambdaFunctions,omitempty" yaml:"lambdaFunctions,omitempty"`

	// Notification configuration to SQS Queue. See below.
	Queues []types.S3_BucketNotificationQueue `json:"queues,omitempty" yaml:"queues,omitempty"`

	// Notification configuration to SNS Topic. See below.
	Topics []types.S3_BucketNotificationTopic `json:"topics,omitempty" yaml:"topics,omitempty"`
}
