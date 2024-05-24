package types

type Ivs_RecordingConfigurationDestinationConfiguration struct {
	// S3 destination configuration where recorded videos will be stored.
	S3 Ivs_RecordingConfigurationDestinationConfigurationS3 `json:"s3,omitempty" yaml:"s3,omitempty"`
}
