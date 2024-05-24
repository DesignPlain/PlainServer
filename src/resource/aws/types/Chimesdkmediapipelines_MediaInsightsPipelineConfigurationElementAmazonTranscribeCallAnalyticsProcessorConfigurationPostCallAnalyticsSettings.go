package types

type Chimesdkmediapipelines_MediaInsightsPipelineConfigurationElementAmazonTranscribeCallAnalyticsProcessorConfigurationPostCallAnalyticsSettings struct {
	// Should output be redacted.
	ContentRedactionOutput string `json:"contentRedactionOutput,omitempty" yaml:"contentRedactionOutput,omitempty"`

	// ARN of the role used by AWS Transcribe to upload your post call analysis.
	DataAccessRoleArn string `json:"dataAccessRoleArn,omitempty" yaml:"dataAccessRoleArn,omitempty"`

	// ID of the KMS key used to encrypt the output.
	OutputEncryptionKmsKeyId string `json:"outputEncryptionKmsKeyId,omitempty" yaml:"outputEncryptionKmsKeyId,omitempty"`

	// The Amazon S3 location where you want your Call Analytics post-call transcription output stored.
	OutputLocation string `json:"outputLocation,omitempty" yaml:"outputLocation,omitempty"`
}
