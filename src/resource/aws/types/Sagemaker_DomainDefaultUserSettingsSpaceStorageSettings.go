package types

type Sagemaker_DomainDefaultUserSettingsSpaceStorageSettings struct {
	// The default EBS storage settings for a private space. See `default_ebs_storage_settings` Block below.
	DefaultEbsStorageSettings Sagemaker_DomainDefaultUserSettingsSpaceStorageSettingsDefaultEbsStorageSettings `json:"defaultEbsStorageSettings,omitempty" yaml:"defaultEbsStorageSettings,omitempty"`
}
