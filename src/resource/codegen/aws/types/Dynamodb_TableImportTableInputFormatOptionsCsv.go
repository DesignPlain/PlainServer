package types

type Dynamodb_TableImportTableInputFormatOptionsCsv struct {
	// The delimiter used for separating items in the CSV file being imported.
	Delimiter string `json:"delimiter,omitempty" yaml:"delimiter,omitempty"`

	// List of the headers used to specify a common header for all source CSV files being imported.
	HeaderLists []string `json:"headerLists,omitempty" yaml:"headerLists,omitempty"`
}
