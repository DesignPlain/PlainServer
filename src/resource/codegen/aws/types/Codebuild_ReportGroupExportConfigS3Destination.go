package types

type Codebuild_ReportGroupExportConfigS3Destination struct {
	// The type of build output artifact to create. Valid values are: `NONE` (default) and `ZIP`.
	Packaging string `json:"packaging,omitempty" yaml:"packaging,omitempty"`

	// The path to the exported report's raw data results.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// The name of the S3 bucket where the raw data of a report are exported.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	/*
	   A boolean value that specifies if the results of a report are encrypted.
	   --Note: the API does not currently allow setting encryption as disabled--
	*/
	EncryptionDisabled bool `json:"encryptionDisabled,omitempty" yaml:"encryptionDisabled,omitempty"`

	// The encryption key for the report's encrypted raw data. The KMS key ARN.
	EncryptionKey string `json:"encryptionKey,omitempty" yaml:"encryptionKey,omitempty"`
}
