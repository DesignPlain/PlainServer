package types

type Imagebuilder_getImageImageScanningConfiguration struct {
	// Configuration block with ECR configuration.
	EcrConfigurations []Imagebuilder_getImageImageScanningConfigurationEcrConfiguration `json:"ecrConfigurations,omitempty" yaml:"ecrConfigurations,omitempty"`

	// Indicates whether Image Builder keeps a snapshot of the vulnerability scans that Amazon Inspector runs against the build instance when you create a new image.
	ImageScanningEnabled bool `json:"imageScanningEnabled,omitempty" yaml:"imageScanningEnabled,omitempty"`
}
