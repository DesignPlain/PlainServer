package types

type Fis_ExperimentTemplateLogConfigurationS3Configuration struct {
	// The name of the destination bucket.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// The bucket prefix.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
