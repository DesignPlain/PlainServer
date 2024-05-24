package types

type Transfer_WorkflowStepDecryptStepDetailsDestinationFileLocationEfsFileLocation struct {
	// The ID of the file system, assigned by Amazon EFS.
	FileSystemId string `json:"fileSystemId,omitempty" yaml:"fileSystemId,omitempty"`

	// The pathname for the folder being used by a workflow.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}
