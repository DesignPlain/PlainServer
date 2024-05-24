package rds

import types "DesignSphere_Server/src/resource/aws/types"

type ExportTask struct {
	// ID of the Amazon Web Services KMS key to use to encrypt the snapshot.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Name of the Amazon S3 bucket to export the snapshot to.
	S3BucketName string `json:"s3BucketName,omitempty" yaml:"s3BucketName,omitempty"`

	// Amazon S3 bucket prefix to use as the file name and path of the exported snapshot.
	S3Prefix string `json:"s3Prefix,omitempty" yaml:"s3Prefix,omitempty"`

	/*
	   Amazon Resource Name (ARN) of the snapshot to export.

	   The following arguments are optional:
	*/
	SourceArn string `json:"sourceArn,omitempty" yaml:"sourceArn,omitempty"`

	//
	Timeouts types.Rds_ExportTaskTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Data to be exported from the snapshot. If this parameter is not provided, all the snapshot data is exported. Valid values are documented in the [AWS StartExportTask API documentation](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_StartExportTask.html#API_StartExportTask_RequestParameters).
	ExportOnlies []string `json:"exportOnlies,omitempty" yaml:"exportOnlies,omitempty"`

	// Unique identifier for the snapshot export task.
	ExportTaskIdentifier string `json:"exportTaskIdentifier,omitempty" yaml:"exportTaskIdentifier,omitempty"`

	// ARN of the IAM role to use for writing to the Amazon S3 bucket.
	IamRoleArn string `json:"iamRoleArn,omitempty" yaml:"iamRoleArn,omitempty"`
}
