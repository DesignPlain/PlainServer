package types

type M2_EnvironmentStorageConfiguration struct {
	//
	Efs M2_EnvironmentStorageConfigurationEfs `json:"efs,omitempty" yaml:"efs,omitempty"`

	//
	Fsx M2_EnvironmentStorageConfigurationFsx `json:"fsx,omitempty" yaml:"fsx,omitempty"`
}
