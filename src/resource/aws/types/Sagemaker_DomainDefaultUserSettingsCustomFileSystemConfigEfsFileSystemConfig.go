package types

type Sagemaker_DomainDefaultUserSettingsCustomFileSystemConfigEfsFileSystemConfig struct {
	// The ID of your Amazon EFS file system.
	FileSystemId string `json:"fileSystemId,omitempty" yaml:"fileSystemId,omitempty"`

	// The path to the file system directory that is accessible in Amazon SageMaker Studio. Permitted users can access only this directory and below.
	FileSystemPath string `json:"fileSystemPath,omitempty" yaml:"fileSystemPath,omitempty"`
}
