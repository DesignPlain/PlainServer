package cloudtrail

import types "libds/aws/types"

type Trail struct {
	// S3 key prefix that follows the name of the bucket you have designated for log file delivery.
	S3KeyPrefix string `json:"s3KeyPrefix,omitempty" yaml:"s3KeyPrefix,omitempty"`

	// Whether log file integrity validation is enabled. Defaults to `false`.
	EnableLogFileValidation bool `json:"enableLogFileValidation,omitempty" yaml:"enableLogFileValidation,omitempty"`

	// Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these. Conflicts with `advanced_event_selector`.
	EventSelectors []types.Cloudtrail_TrailEventSelector `json:"eventSelectors,omitempty" yaml:"eventSelectors,omitempty"`

	// Whether the trail is publishing events from global services such as IAM to the log files. Defaults to `true`.
	IncludeGlobalServiceEvents bool `json:"includeGlobalServiceEvents,omitempty" yaml:"includeGlobalServiceEvents,omitempty"`

	// Configuration block for identifying unusual operational activity. See details below.
	InsightSelectors []types.Cloudtrail_TrailInsightSelector `json:"insightSelectors,omitempty" yaml:"insightSelectors,omitempty"`

	// Name of the trail.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies an advanced event selector for enabling data event logging. Fields documented below. Conflicts with `event_selector`.
	AdvancedEventSelectors []types.Cloudtrail_TrailAdvancedEventSelector `json:"advancedEventSelectors,omitempty" yaml:"advancedEventSelectors,omitempty"`

	// Role for the CloudWatch Logs endpoint to assume to write to a user’s log group.
	CloudWatchLogsRoleArn string `json:"cloudWatchLogsRoleArn,omitempty" yaml:"cloudWatchLogsRoleArn,omitempty"`

	// Whether the trail is created in the current region or in all regions. Defaults to `false`.
	IsMultiRegionTrail bool `json:"isMultiRegionTrail,omitempty" yaml:"isMultiRegionTrail,omitempty"`

	/*
	   Name of the S3 bucket designated for publishing log files.

	   The following arguments are optional:
	*/
	S3BucketName string `json:"s3BucketName,omitempty" yaml:"s3BucketName,omitempty"`

	// Log group name using an ARN that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	CloudWatchLogsGroupArn string `json:"cloudWatchLogsGroupArn,omitempty" yaml:"cloudWatchLogsGroupArn,omitempty"`

	// Enables logging for the trail. Defaults to `true`. Setting this to `false` will pause logging.
	EnableLogging bool `json:"enableLogging,omitempty" yaml:"enableLogging,omitempty"`

	// Whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
	IsOrganizationTrail bool `json:"isOrganizationTrail,omitempty" yaml:"isOrganizationTrail,omitempty"`

	// Name of the Amazon SNS topic defined for notification of log file delivery.
	SnsTopicName string `json:"snsTopicName,omitempty" yaml:"snsTopicName,omitempty"`

	// KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Map of tags to assign to the trail. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
