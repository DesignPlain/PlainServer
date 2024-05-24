package types

type Sagemaker_UserProfileUserSettingsCustomFileSystemConfig struct {
	// The default EBS storage settings for a private space. See EFS File System Config below.
	EfsFileSystemConfigs []Sagemaker_UserProfileUserSettingsCustomFileSystemConfigEfsFileSystemConfig `json:"efsFileSystemConfigs,omitempty" yaml:"efsFileSystemConfigs,omitempty"`
}
