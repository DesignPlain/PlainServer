package types

type Bigquery_TableTableConstraintsForeignKeyReferencedTable struct {
	// The ID of the dataset containing this table.
	DatasetId string `json:"datasetId,omitempty" yaml:"datasetId,omitempty"`

	// The ID of the project containing this table.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`

	/*
	   The ID of the table. The ID must contain only
	   letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum
	   length is 1,024 characters. Certain operations allow suffixing of
	   the table ID with a partition decorator, such as
	   sample_table$20190123.
	*/
	TableId string `json:"tableId,omitempty" yaml:"tableId,omitempty"`
}
