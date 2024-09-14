package types

type Transfer_WorkflowStepDecryptStepDetailsDestinationFileLocation struct {
	// Specifies the details for the S3 file being copied.
	S3FileLocation Transfer_WorkflowStepDecryptStepDetailsDestinationFileLocationS3FileLocation `json:"s3FileLocation,omitempty" yaml:"s3FileLocation,omitempty"`

	// Specifies the details for the EFS file being copied.
	EfsFileLocation Transfer_WorkflowStepDecryptStepDetailsDestinationFileLocationEfsFileLocation `json:"efsFileLocation,omitempty" yaml:"efsFileLocation,omitempty"`
}
