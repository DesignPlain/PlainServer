package types

type Sagemaker_UserProfileUserSettingsSpaceStorageSettingsDefaultEbsStorageSettings struct {
	// The default size of the EBS storage volume for a private space.
	DefaultEbsVolumeSizeInGb int `json:"defaultEbsVolumeSizeInGb,omitempty" yaml:"defaultEbsVolumeSizeInGb,omitempty"`

	// The maximum size of the EBS storage volume for a private space.
	MaximumEbsVolumeSizeInGb int `json:"maximumEbsVolumeSizeInGb,omitempty" yaml:"maximumEbsVolumeSizeInGb,omitempty"`
}
