package types

type Quicksight_DataSetPhysicalTableMapS3Source struct {
	// ARN of the data source.
	DataSourceArn string `json:"dataSourceArn,omitempty" yaml:"dataSourceArn,omitempty"`

	// Column schema of the table. See input_columns.
	InputColumns []Quicksight_DataSetPhysicalTableMapS3SourceInputColumn `json:"inputColumns,omitempty" yaml:"inputColumns,omitempty"`

	// Information about the format for the S3 source file or files. See upload_settings.
	UploadSettings Quicksight_DataSetPhysicalTableMapS3SourceUploadSettings `json:"uploadSettings,omitempty" yaml:"uploadSettings,omitempty"`
}
