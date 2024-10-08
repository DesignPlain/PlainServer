package types

type Quicksight_DataSourceParametersTwitter struct {
	// The Twitter query to retrieve the data.
	Query string `json:"query,omitempty" yaml:"query,omitempty"`

	// The maximum number of rows to query.
	MaxRows int `json:"maxRows,omitempty" yaml:"maxRows,omitempty"`
}
