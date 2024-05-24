package types

type Macie2_ClassificationExportConfigurationS3Destination struct {
	// The Amazon S3 bucket name in which Amazon Macie exports the data classification results.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// The object key for the bucket in which Amazon Macie exports the data classification results.
	KeyPrefix string `json:"keyPrefix,omitempty" yaml:"keyPrefix,omitempty"`

	/*
	   Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.

	   Additional information can be found in the [Storing and retaining sensitive data discovery results with Amazon Macie for AWS Macie documentation](https://docs.aws.amazon.com/macie/latest/user/discovery-results-repository-s3.html).
	*/
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`
}
