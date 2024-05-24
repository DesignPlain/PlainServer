package types

type Quicksight_DataSetPhysicalTableMap struct {
	// A physical table type built from the results of the custom SQL query. See custom_sql.
	CustomSql Quicksight_DataSetPhysicalTableMapCustomSql `json:"customSql,omitempty" yaml:"customSql,omitempty"`

	// Key of the physical table map.
	PhysicalTableMapId string `json:"physicalTableMapId,omitempty" yaml:"physicalTableMapId,omitempty"`

	// A physical table type for relational data sources. See relational_table.
	RelationalTable Quicksight_DataSetPhysicalTableMapRelationalTable `json:"relationalTable,omitempty" yaml:"relationalTable,omitempty"`

	// A physical table type for as S3 data source. See s3_source.
	S3Source Quicksight_DataSetPhysicalTableMapS3Source `json:"s3Source,omitempty" yaml:"s3Source,omitempty"`
}
