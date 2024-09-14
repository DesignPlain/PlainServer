package types

type Transfer_WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocation struct {
	// Specifies the details for the EFS file being copied.
	EfsFileLocation Transfer_WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationEfsFileLocation `json:"efsFileLocation,omitempty" yaml:"efsFileLocation,omitempty"`

	// Specifies the details for the S3 file being copied.
	S3FileLocation Transfer_WorkflowOnExceptionStepDecryptStepDetailsDestinationFileLocationS3FileLocation `json:"s3FileLocation,omitempty" yaml:"s3FileLocation,omitempty"`
}
