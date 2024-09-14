package cleanrooms

import types "libds/aws/types"

type ConfiguredTable struct {
	// The columns of the references table which will be included in the configured table.
	AllowedColumns []string `json:"allowedColumns,omitempty" yaml:"allowedColumns,omitempty"`

	// The analysis method for the configured table. The only valid value is currently `DIRECT_QUERY`.
	AnalysisMethod string `json:"analysisMethod,omitempty" yaml:"analysisMethod,omitempty"`

	// A description for the configured table.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the configured table.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   A reference to the AWS Glue table which will be used to create the configured table.
	   - `table_reference.database_name` - (Required - Forces new resource) - The name of the AWS Glue database which contains the table.
	   - `table_reference.table_name` - (Required - Forces new resource) - The name of the AWS Glue table which will be used to create the configured table.
	*/
	TableReference types.Cleanrooms_ConfiguredTableTableReference `json:"tableReference,omitempty" yaml:"tableReference,omitempty"`

	// Key value pairs which tag the configured table.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
