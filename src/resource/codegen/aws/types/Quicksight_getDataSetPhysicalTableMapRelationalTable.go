package types

type Quicksight_getDataSetPhysicalTableMapRelationalTable struct {
	//
	DataSourceArn string `json:"dataSourceArn,omitempty" yaml:"dataSourceArn,omitempty"`

	//
	InputColumns []Quicksight_getDataSetPhysicalTableMapRelationalTableInputColumn `json:"inputColumns,omitempty" yaml:"inputColumns,omitempty"`

	//
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`

	//
	Catalog string `json:"catalog,omitempty" yaml:"catalog,omitempty"`
}
