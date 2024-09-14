package types

type Sagemaker_WorkteamWorkerAccessConfiguration struct {
	// Defines any Amazon S3 resource constraints. see S3 Presign details below.
	S3Presign Sagemaker_WorkteamWorkerAccessConfigurationS3Presign `json:"s3Presign,omitempty" yaml:"s3Presign,omitempty"`
}
