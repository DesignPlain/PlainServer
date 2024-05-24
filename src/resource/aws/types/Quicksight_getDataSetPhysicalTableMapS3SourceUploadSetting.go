package types

type Quicksight_getDataSetPhysicalTableMapS3SourceUploadSetting struct {
	//
	Format string `json:"format,omitempty" yaml:"format,omitempty"`

	//
	StartFromRow int `json:"startFromRow,omitempty" yaml:"startFromRow,omitempty"`

	//
	TextQualifier string `json:"textQualifier,omitempty" yaml:"textQualifier,omitempty"`

	//
	ContainsHeader bool `json:"containsHeader,omitempty" yaml:"containsHeader,omitempty"`

	//
	Delimiter string `json:"delimiter,omitempty" yaml:"delimiter,omitempty"`
}
