package types

type Glue_ClassifierCsvClassifier struct {
	// Enables the processing of files that contain only one column.
	AllowSingleColumn bool `json:"allowSingleColumn,omitempty" yaml:"allowSingleColumn,omitempty"`

	// Indicates whether the CSV file contains a header. This can be one of "ABSENT", "PRESENT", or "UNKNOWN".
	ContainsHeader string `json:"containsHeader,omitempty" yaml:"containsHeader,omitempty"`

	// Enables the custom datatype to be configured.
	CustomDatatypeConfigured bool `json:"customDatatypeConfigured,omitempty" yaml:"customDatatypeConfigured,omitempty"`

	// A list of supported custom data Valid values are `BINARY`, `BOOLEAN`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `INT`, `LONG`, `SHORT`, `STRING`, `TIMESTAMP`.
	CustomDatatypes []string `json:"customDatatypes,omitempty" yaml:"customDatatypes,omitempty"`

	// The delimiter used in the Csv to separate columns.
	Delimiter string `json:"delimiter,omitempty" yaml:"delimiter,omitempty"`

	// Specifies whether to trim column values.
	DisableValueTrimming bool `json:"disableValueTrimming,omitempty" yaml:"disableValueTrimming,omitempty"`

	// A list of strings representing column names.
	Headers []string `json:"headers,omitempty" yaml:"headers,omitempty"`

	// A custom symbol to denote what combines content into a single column value. It must be different from the column delimiter.
	QuoteSymbol string `json:"quoteSymbol,omitempty" yaml:"quoteSymbol,omitempty"`

	//
	Serde string `json:"serde,omitempty" yaml:"serde,omitempty"`
}
