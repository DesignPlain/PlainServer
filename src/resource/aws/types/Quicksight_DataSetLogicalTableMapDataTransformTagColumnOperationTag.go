package types

type Quicksight_DataSetLogicalTableMapDataTransformTagColumnOperationTag struct {
	// A description for a column. See column_description.
	ColumnDescription Quicksight_DataSetLogicalTableMapDataTransformTagColumnOperationTagColumnDescription `json:"columnDescription,omitempty" yaml:"columnDescription,omitempty"`

	// A geospatial role for a column. Valid values are `COUNTRY`, `STATE`, `COUNTY`, `CITY`, `POSTCODE`, `LONGITUDE`, and `LATITUDE`.
	ColumnGeographicRole string `json:"columnGeographicRole,omitempty" yaml:"columnGeographicRole,omitempty"`
}
