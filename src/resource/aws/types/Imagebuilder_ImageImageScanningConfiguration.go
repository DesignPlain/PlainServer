package types

type Imagebuilder_ImageImageScanningConfiguration struct {
	// Configuration block with ECR configuration. Detailed below.
	EcrConfiguration Imagebuilder_ImageImageScanningConfigurationEcrConfiguration `json:"ecrConfiguration,omitempty" yaml:"ecrConfiguration,omitempty"`

	// Indicates whether Image Builder keeps a snapshot of the vulnerability scans that Amazon Inspector runs against the build instance when you create a new image. Defaults to `false`.
	ImageScanningEnabled bool `json:"imageScanningEnabled,omitempty" yaml:"imageScanningEnabled,omitempty"`
}
