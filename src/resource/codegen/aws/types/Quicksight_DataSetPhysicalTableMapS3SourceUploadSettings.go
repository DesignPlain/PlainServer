package types

type Quicksight_DataSetPhysicalTableMapS3SourceUploadSettings struct {
	// Delimiter between values in the file.
	Delimiter string `json:"delimiter,omitempty" yaml:"delimiter,omitempty"`

	// File format. Valid values are `CSV`, `TSV`, `CLF`, `ELF`, `XLSX`, and `JSON`.
	Format string `json:"format,omitempty" yaml:"format,omitempty"`

	// A row number to start reading data from.
	StartFromRow int `json:"startFromRow,omitempty" yaml:"startFromRow,omitempty"`

	// Text qualifier. Valid values are `DOUBLE_QUOTE` and `SINGLE_QUOTE`.
	TextQualifier string `json:"textQualifier,omitempty" yaml:"textQualifier,omitempty"`

	// Whether the file has a header row, or the files each have a header row.
	ContainsHeader bool `json:"containsHeader,omitempty" yaml:"containsHeader,omitempty"`
}
