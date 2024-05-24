package types

type Transfer_WorkflowStepDecryptStepDetailsDestinationFileLocationS3FileLocation struct {
	// The name assigned to the tag that you create.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Specifies the S3 bucket for the customer input file.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`
}
