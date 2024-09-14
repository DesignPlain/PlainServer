package types

type Quicksight_getDataSetColumnLevelPermissionRule struct {
	//
	ColumnNames []string `json:"columnNames,omitempty" yaml:"columnNames,omitempty"`

	//
	Principals []string `json:"principals,omitempty" yaml:"principals,omitempty"`
}
