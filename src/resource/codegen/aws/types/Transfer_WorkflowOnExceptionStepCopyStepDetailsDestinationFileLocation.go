package types

type Transfer_WorkflowOnExceptionStepCopyStepDetailsDestinationFileLocation struct {
	// Specifies the details for the EFS file being copied.
	EfsFileLocation Transfer_WorkflowOnExceptionStepCopyStepDetailsDestinationFileLocationEfsFileLocation `json:"efsFileLocation,omitempty" yaml:"efsFileLocation,omitempty"`

	// Specifies the details for the S3 file being copied.
	S3FileLocation Transfer_WorkflowOnExceptionStepCopyStepDetailsDestinationFileLocationS3FileLocation `json:"s3FileLocation,omitempty" yaml:"s3FileLocation,omitempty"`
}
