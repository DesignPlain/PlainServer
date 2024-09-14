package types

type Quicksight_DataSourceParametersS3 struct {
	// An object containing the S3 location of the S3 manifest file.
	ManifestFileLocation Quicksight_DataSourceParametersS3ManifestFileLocation `json:"manifestFileLocation,omitempty" yaml:"manifestFileLocation,omitempty"`
}
