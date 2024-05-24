package types

type Quicksight_getDataSetPhysicalTableMapRelationalTable struct {
	//
	Catalog string `json:"catalog,omitempty" yaml:"catalog,omitempty"`

	//
	DataSourceArn string `json:"dataSourceArn,omitempty" yaml:"dataSourceArn,omitempty"`

	//
	InputColumns []Quicksight_getDataSetPhysicalTableMapRelationalTableInputColumn `json:"inputColumns,omitempty" yaml:"inputColumns,omitempty"`

	//
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Schema string `json:"schema,omitempty" yaml:"schema,omitempty"`
}
