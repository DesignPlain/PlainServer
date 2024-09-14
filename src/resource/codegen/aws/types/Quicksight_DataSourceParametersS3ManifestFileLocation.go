package types

type Quicksight_DataSourceParametersS3ManifestFileLocation struct {
	// The name of the bucket that contains the manifest file.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// The key of the manifest file within the bucket.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
