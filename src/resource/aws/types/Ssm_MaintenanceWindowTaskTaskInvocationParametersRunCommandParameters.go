package types

type Ssm_MaintenanceWindowTaskTaskInvocationParametersRunCommandParameters struct {
	// SHA-256 or SHA-1. SHA-1 hashes have been deprecated. Valid values: `Sha256` and `Sha1`
	DocumentHashType string `json:"documentHashType,omitempty" yaml:"documentHashType,omitempty"`

	// The name of the Amazon S3 bucket.
	OutputS3Bucket string `json:"outputS3Bucket,omitempty" yaml:"outputS3Bucket,omitempty"`

	// The Amazon S3 bucket subfolder.
	OutputS3KeyPrefix string `json:"outputS3KeyPrefix,omitempty" yaml:"outputS3KeyPrefix,omitempty"`

	// The parameters for the RUN_COMMAND task execution. Documented below.
	Parameters []Ssm_MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Configuration options for sending command output to CloudWatch Logs. Documented below.
	CloudwatchConfig Ssm_MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersCloudwatchConfig `json:"cloudwatchConfig,omitempty" yaml:"cloudwatchConfig,omitempty"`

	// Information about the command(s) to execute.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// The SHA-256 or SHA-1 hash created by the system when the document was created. SHA-1 hashes have been deprecated.
	DocumentHash string `json:"documentHash,omitempty" yaml:"documentHash,omitempty"`

	// The version of an Automation document to use during task execution.
	DocumentVersion string `json:"documentVersion,omitempty" yaml:"documentVersion,omitempty"`

	// Configurations for sending notifications about command status changes on a per-instance basis. Documented below.
	NotificationConfig Ssm_MaintenanceWindowTaskTaskInvocationParametersRunCommandParametersNotificationConfig `json:"notificationConfig,omitempty" yaml:"notificationConfig,omitempty"`

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) service role to use to publish Amazon Simple Notification Service (Amazon SNS) notifications for maintenance window Run Command tasks.
	ServiceRoleArn string `json:"serviceRoleArn,omitempty" yaml:"serviceRoleArn,omitempty"`

	// If this time is reached and the command has not already started executing, it doesn't run.
	TimeoutSeconds int `json:"timeoutSeconds,omitempty" yaml:"timeoutSeconds,omitempty"`
}
