package types

type Glue_ClassifierCsvClassifier struct {
	// Enables the custom datatype to be configured.
	CustomDatatypeConfigured bool `json:"customDatatypeConfigured,omitempty" yaml:"customDatatypeConfigured,omitempty"`

	// Specifies whether to trim column values.
	DisableValueTrimming bool `json:"disableValueTrimming,omitempty" yaml:"disableValueTrimming,omitempty"`

	// A list of strings representing column names.
	Headers []string `json:"headers,omitempty" yaml:"headers,omitempty"`

	// Enables the processing of files that contain only one column.
	AllowSingleColumn bool `json:"allowSingleColumn,omitempty" yaml:"allowSingleColumn,omitempty"`

	// Indicates whether the CSV file contains a header. This can be one of "ABSENT", "PRESENT", or "UNKNOWN".
	ContainsHeader string `json:"containsHeader,omitempty" yaml:"containsHeader,omitempty"`

	// A list of supported custom datatypes. Valid values are `BINARY`, `BOOLEAN`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `INT`, `LONG`, `SHORT`, `STRING`, `TIMESTAMP`.
	CustomDatatypes []string `json:"customDatatypes,omitempty" yaml:"customDatatypes,omitempty"`

	// The delimiter used in the CSV to separate columns.
	Delimiter string `json:"delimiter,omitempty" yaml:"delimiter,omitempty"`

	// A custom symbol to denote what combines content into a single column value. It must be different from the column delimiter.
	QuoteSymbol string `json:"quoteSymbol,omitempty" yaml:"quoteSymbol,omitempty"`

	// The SerDe for processing CSV. Valid values are `OpenCSVSerDe`, `LazySimpleSerDe`, `None`.
	Serde string `json:"serde,omitempty" yaml:"serde,omitempty"`
}
