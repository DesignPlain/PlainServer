package types

type Bigqueryanalyticshub_ListingRestrictedExportConfig struct {
	// If true, enable restricted export.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// If true, restrict export of query result derived from restricted linked dataset table.
	RestrictQueryResult bool `json:"restrictQueryResult,omitempty" yaml:"restrictQueryResult,omitempty"`
}
