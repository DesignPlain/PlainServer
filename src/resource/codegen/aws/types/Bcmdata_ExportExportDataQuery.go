package types

type Bcmdata_ExportExportDataQuery struct {
	// Query statement.
	QueryStatement string `json:"queryStatement,omitempty" yaml:"queryStatement,omitempty"`

	// Table configuration.
	TableConfigurations map[string]map[string]string `json:"tableConfigurations,omitempty" yaml:"tableConfigurations,omitempty"`
}
