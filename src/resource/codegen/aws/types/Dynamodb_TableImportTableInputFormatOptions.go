package types

type Dynamodb_TableImportTableInputFormatOptions struct {
	// This block contains the processing options for the CSV file being imported:
	Csv Dynamodb_TableImportTableInputFormatOptionsCsv `json:"csv,omitempty" yaml:"csv,omitempty"`
}
