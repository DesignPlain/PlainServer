package types

type Sagemaker_SpaceSpaceSettingsKernelGatewayAppSettingsCustomImage struct {
	// The version number of the Custom Image.
	ImageVersionNumber int `json:"imageVersionNumber,omitempty" yaml:"imageVersionNumber,omitempty"`

	// The name of the App Image Config.
	AppImageConfigName string `json:"appImageConfigName,omitempty" yaml:"appImageConfigName,omitempty"`

	// The name of the Custom Image.
	ImageName string `json:"imageName,omitempty" yaml:"imageName,omitempty"`
}
