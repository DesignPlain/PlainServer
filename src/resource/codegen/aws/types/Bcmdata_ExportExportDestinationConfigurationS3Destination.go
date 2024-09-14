package types

type Bcmdata_ExportExportDestinationConfigurationS3Destination struct {
	// S3 bucket region.
	S3Region string `json:"s3Region,omitempty" yaml:"s3Region,omitempty"`

	// Name of the Amazon S3 bucket used as the destination of a data export file.
	S3Bucket string `json:"s3Bucket,omitempty" yaml:"s3Bucket,omitempty"`

	// Output configuration for the data export. See the `s3_output_configurations` argument reference below.
	S3OutputConfigurations []Bcmdata_ExportExportDestinationConfigurationS3DestinationS3OutputConfiguration `json:"s3OutputConfigurations,omitempty" yaml:"s3OutputConfigurations,omitempty"`

	// S3 path prefix you want prepended to the name of your data export.
	S3Prefix string `json:"s3Prefix,omitempty" yaml:"s3Prefix,omitempty"`
}
