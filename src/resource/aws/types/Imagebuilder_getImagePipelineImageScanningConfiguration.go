package types

type Imagebuilder_getImagePipelineImageScanningConfiguration struct {
	// List if an object with ecr configuration for image scanning
	EcrConfigurations []Imagebuilder_getImagePipelineImageScanningConfigurationEcrConfiguration `json:"ecrConfigurations,omitempty" yaml:"ecrConfigurations,omitempty"`

	// Whether image scanning is enabled.
	ImageScanningEnabled bool `json:"imageScanningEnabled,omitempty" yaml:"imageScanningEnabled,omitempty"`
}
