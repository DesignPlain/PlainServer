package types

type Sagemaker_SpaceSpaceSettingsCustomFileSystem struct {
	// A custom file system in Amazon EFS. See `efs_file_system` Block below.
	EfsFileSystem Sagemaker_SpaceSpaceSettingsCustomFileSystemEfsFileSystem `json:"efsFileSystem,omitempty" yaml:"efsFileSystem,omitempty"`
}
