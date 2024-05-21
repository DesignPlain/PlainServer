package types

type Bigquery_TableExternalDataConfigurationCsvOptions struct {
	/*
	   Indicates if BigQuery should accept rows
	   that are missing trailing optional columns.
	*/
	AllowJaggedRows bool `json:"allowJaggedRows,omitempty" yaml:"allowJaggedRows,omitempty"`

	/*
	   Indicates if BigQuery should allow
	   quoted data sections that contain newline characters in a CSV file.
	   The default value is false.
	*/
	AllowQuotedNewlines bool `json:"allowQuotedNewlines,omitempty" yaml:"allowQuotedNewlines,omitempty"`

	/*
	   The character encoding of the data. The supported
	   values are UTF-8 or ISO-8859-1.
	*/
	Encoding string `json:"encoding,omitempty" yaml:"encoding,omitempty"`

	// The separator for fields in a CSV file.
	FieldDelimiter string `json:"fieldDelimiter,omitempty" yaml:"fieldDelimiter,omitempty"`

	/*
	   The value that is used to quote data sections in a
	   CSV file. If your data does not contain quoted sections, set the
	   property value to an empty string. If your data contains quoted newline
	   characters, you must also set the `allow_quoted_newlines` property to true.
	   The API-side default is `"`, specified in the provider escaped as `\"`. Due to
	   limitations with default values, this value is required to be
	   explicitly set.
	*/
	Quote string `json:"quote,omitempty" yaml:"quote,omitempty"`

	/*
	   The number of rows at the top of a CSV
	   file that BigQuery will skip when reading the data.
	*/
	SkipLeadingRows int `json:"skipLeadingRows,omitempty" yaml:"skipLeadingRows,omitempty"`
}
