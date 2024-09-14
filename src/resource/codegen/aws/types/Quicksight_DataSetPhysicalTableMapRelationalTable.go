package types

type Quicksight_DataSetPhysicalTableMapRelationalTable struct {
	// Catalog associated with the table.
	Catalog string `json:"catalog,omitempty" yaml:"catalog,omitempty"`

	// ARN of the data source.
	DataSourceArn string `json:"dataSourceArn,omitempty" yaml:"dataSourceArn,omitempty"`

	// Column schema of the table. See input_columns.
	InputColumns []Quicksight_DataSetPhysicalTableMapRelationalTableInputColumn `json:"inputColumns,omitempty" yaml:"inputColumns,omitempty"`

	// Name of the relational table.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Schema name. This name applies to certain relational database engines.
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`
}
