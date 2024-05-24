package types

type Sagemaker_DomainDefaultUserSettingsCustomFileSystemConfig struct {
	// The default EBS storage settings for a private space. See `efs_file_system_config` Block below.
	EfsFileSystemConfig Sagemaker_DomainDefaultUserSettingsCustomFileSystemConfigEfsFileSystemConfig `json:"efsFileSystemConfig,omitempty" yaml:"efsFileSystemConfig,omitempty"`
}
