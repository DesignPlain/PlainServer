package types

type Transfer_WorkflowStepCopyStepDetailsDestinationFileLocation struct {
	// Specifies the details for the EFS file being copied.
	EfsFileLocation Transfer_WorkflowStepCopyStepDetailsDestinationFileLocationEfsFileLocation `json:"efsFileLocation,omitempty" yaml:"efsFileLocation,omitempty"`

	// Specifies the details for the S3 file being copied.
	S3FileLocation Transfer_WorkflowStepCopyStepDetailsDestinationFileLocationS3FileLocation `json:"s3FileLocation,omitempty" yaml:"s3FileLocation,omitempty"`
}
