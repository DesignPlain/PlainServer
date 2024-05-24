package types

type Quicksight_getDataSetPhysicalTableMapS3Source struct {
	//
	DataSourceArn string `json:"dataSourceArn,omitempty" yaml:"dataSourceArn,omitempty"`

	//
	InputColumns []Quicksight_getDataSetPhysicalTableMapS3SourceInputColumn `json:"inputColumns,omitempty" yaml:"inputColumns,omitempty"`

	//
	UploadSettings []Quicksight_getDataSetPhysicalTableMapS3SourceUploadSetting `json:"uploadSettings,omitempty" yaml:"uploadSettings,omitempty"`
}
