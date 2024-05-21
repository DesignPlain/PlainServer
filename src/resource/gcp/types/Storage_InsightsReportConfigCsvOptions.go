package types

type Storage_InsightsReportConfigCsvOptions struct {
	/*
	   The boolean that indicates whether or not headers are included in the inventory report CSV file.

	   - - -
	*/
	HeaderRequired bool `json:"headerRequired,omitempty" yaml:"headerRequired,omitempty"`

	// The character used to separate the records in the inventory report CSV file.
	RecordSeparator string `json:"recordSeparator,omitempty" yaml:"recordSeparator,omitempty"`

	// The delimiter used to separate the fields in the inventory report CSV file.
	Delimiter string `json:"delimiter,omitempty" yaml:"delimiter,omitempty"`
}
