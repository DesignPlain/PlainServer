package types

type Bedrock_AgentAgentActionGroupApiSchemaS3 struct {
	// Name of the S3 bucket.
	S3BucketName string `json:"s3BucketName,omitempty" yaml:"s3BucketName,omitempty"`

	// S3 object key containing the resource.
	S3ObjectKey string `json:"s3ObjectKey,omitempty" yaml:"s3ObjectKey,omitempty"`
}
