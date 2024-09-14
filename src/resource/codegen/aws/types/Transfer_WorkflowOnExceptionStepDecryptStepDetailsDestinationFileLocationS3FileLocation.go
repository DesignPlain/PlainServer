package types

type Transfer_WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationS3FileLocation struct {
	// The name assigned to the file when it was created in S3. You use the object key to retrieve the object.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Specifies the S3 bucket for the customer input file.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`
}
