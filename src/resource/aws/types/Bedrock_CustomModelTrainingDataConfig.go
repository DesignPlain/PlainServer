package types

type Bedrock_CustomModelTrainingDataConfig struct {
	// The S3 URI where the validation data is stored.
	S3Uri string `json:"s3Uri,omitempty" yaml:"s3Uri,omitempty"`
}
