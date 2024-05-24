package cfg

import types "DesignSphere_Server/src/resource/aws/types"

type DeliveryChannel struct {
	// The ARN of the AWS KMS key used to encrypt objects delivered by AWS Config. Must belong to the same Region as the destination S3 bucket.
	S3KmsKeyArn string `json:"s3KmsKeyArn,omitempty" yaml:"s3KmsKeyArn,omitempty"`

	// Options for how AWS Config delivers configuration snapshots. See below
	SnapshotDeliveryProperties types.Cfg_DeliveryChannelSnapshotDeliveryProperties `json:"snapshotDeliveryProperties,omitempty" yaml:"snapshotDeliveryProperties,omitempty"`

	// The ARN of the SNS topic that AWS Config delivers notifications to.
	SnsTopicArn string `json:"snsTopicArn,omitempty" yaml:"snsTopicArn,omitempty"`

	// The name of the delivery channel. Defaults to `default`. Changing it recreates the resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The name of the S3 bucket used to store the configuration history.
	S3BucketName string `json:"s3BucketName,omitempty" yaml:"s3BucketName,omitempty"`

	// The prefix for the specified S3 bucket.
	S3KeyPrefix string `json:"s3KeyPrefix,omitempty" yaml:"s3KeyPrefix,omitempty"`
}
