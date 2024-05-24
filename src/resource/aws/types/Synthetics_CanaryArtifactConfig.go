package types

type Synthetics_CanaryArtifactConfig struct {
	// Configuration of the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3. See S3 Encryption.
	S3Encryption Synthetics_CanaryArtifactConfigS3Encryption `json:"s3Encryption,omitempty" yaml:"s3Encryption,omitempty"`
}
