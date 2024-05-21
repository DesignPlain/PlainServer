package types

type Bigquery_TableExternalDataConfigurationGoogleSheetsOptions struct {
	/*
	   Range of a sheet to query from. Only used when
	   non-empty. At least one of `range` or `skip_leading_rows` must be set.
	   Typical format: "sheet_name!top_left_cell_id:bottom_right_cell_id"
	   For example: "sheet1!A1:B20"
	*/
	Range string `json:"range,omitempty" yaml:"range,omitempty"`

	/*
	   The number of rows at the top of the sheet
	   that BigQuery will skip when reading the data. At least one of `range` or
	   `skip_leading_rows` must be set.
	*/
	SkipLeadingRows int `json:"skipLeadingRows,omitempty" yaml:"skipLeadingRows,omitempty"`
}
