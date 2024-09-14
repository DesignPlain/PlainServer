package types

type Bcmdata_ExportExportDestinationConfigurationS3DestinationS3OutputConfiguration struct {
	// Output type for the data export. Valid value `CUSTOM`.
	OutputType string `json:"outputType,omitempty" yaml:"outputType,omitempty"`

	// The rule to follow when generating a version of the data export file. You have the choice to overwrite the previous version or to be delivered in addition to the previous versions. Overwriting exports can save on Amazon S3 storage costs. Creating new export versions allows you to track the changes in cost and usage data over time. Valid values `CREATE_NEW_REPORT` or `OVERWRITE_REPORT`.
	Overwrite string `json:"overwrite,omitempty" yaml:"overwrite,omitempty"`

	// Compression type for the data export. Valid values `GZIP`, `PARQUET`.
	Compression string `json:"compression,omitempty" yaml:"compression,omitempty"`

	// File format for the data export. Valid values `TEXT_OR_CSV` or `PARQUET`.
	Format string `json:"format,omitempty" yaml:"format,omitempty"`
}
