package types

type Imagebuilder_ImagePipelineImageScanningConfiguration struct {
	// Configuration block with ECR configuration for image scanning. Detailed below.
	EcrConfiguration Imagebuilder_ImagePipelineImageScanningConfigurationEcrConfiguration `json:"ecrConfiguration,omitempty" yaml:"ecrConfiguration,omitempty"`

	// Whether image scans are enabled. Defaults to `false`.
	ImageScanningEnabled bool `json:"imageScanningEnabled,omitempty" yaml:"imageScanningEnabled,omitempty"`
}
