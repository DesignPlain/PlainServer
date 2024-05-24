package types

type Dms_EndpointRedshiftSettings struct {
	// ARN or Id of KMS Key to use when `encryption_mode` is `SSE_KMS`.
	ServerSideEncryptionKmsKeyId string `json:"serverSideEncryptionKmsKeyId,omitempty" yaml:"serverSideEncryptionKmsKeyId,omitempty"`

	// Amazon Resource Name (ARN) of the IAM Role with permissions to read from or write to the S3 Bucket for intermediate storage.
	ServiceAccessRoleArn string `json:"serviceAccessRoleArn,omitempty" yaml:"serviceAccessRoleArn,omitempty"`

	// Custom S3 Bucket Object prefix for intermediate storage.
	BucketFolder string `json:"bucketFolder,omitempty" yaml:"bucketFolder,omitempty"`

	// Custom S3 Bucket name for intermediate storage.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// The server-side encryption mode that you want to encrypt your intermediate .csv object files copied to S3. Defaults to `SSE_S3`. Valid values are `SSE_S3` and `SSE_KMS`.
	EncryptionMode string `json:"encryptionMode,omitempty" yaml:"encryptionMode,omitempty"`
}
