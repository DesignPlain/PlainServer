package types

type Quicksight_DataSetColumnLevelPermissionRule struct {
	// An array of ARNs for Amazon QuickSight users or groups.
	Principals []string `json:"principals,omitempty" yaml:"principals,omitempty"`

	// An array of column names.
	ColumnNames []string `json:"columnNames,omitempty" yaml:"columnNames,omitempty"`
}
