package types

type S3control_StorageLensConfigurationStorageLensConfigurationExclude struct {
	// List of S3 bucket ARNs.
	Buckets []string `json:"buckets,omitempty" yaml:"buckets,omitempty"`

	// List of AWS Regions.
	Regions []string `json:"regions,omitempty" yaml:"regions,omitempty"`
}
