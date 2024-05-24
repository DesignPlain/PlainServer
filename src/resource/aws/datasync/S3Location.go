package datasync

import types "DesignSphere_Server/src/resource/aws/types"

type S3Location struct {
	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns []string `json:"agentArns,omitempty" yaml:"agentArns,omitempty"`

	// Amazon Resource Name (ARN) of the S3 Bucket.
	S3BucketArn string `json:"s3BucketArn,omitempty" yaml:"s3BucketArn,omitempty"`

	// Configuration block containing information for connecting to S3.
	S3Config types.Datasync_S3LocationS3Config `json:"s3Config,omitempty" yaml:"s3Config,omitempty"`

	// The Amazon S3 storage class that you want to store your files in when this location is used as a task destination. [Valid values](https://docs.aws.amazon.com/datasync/latest/userguide/create-s3-location.html#using-storage-classes)
	S3StorageClass string `json:"s3StorageClass,omitempty" yaml:"s3StorageClass,omitempty"`

	// Prefix to perform actions as source or destination.
	Subdirectory string `json:"subdirectory,omitempty" yaml:"subdirectory,omitempty"`

	// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
