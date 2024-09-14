package types

type Quicksight_getDataSetPhysicalTableMap struct {
	//
	CustomSqls []Quicksight_getDataSetPhysicalTableMapCustomSql `json:"customSqls,omitempty" yaml:"customSqls,omitempty"`

	//
	PhysicalTableMapId string `json:"physicalTableMapId,omitempty" yaml:"physicalTableMapId,omitempty"`

	//
	RelationalTables []Quicksight_getDataSetPhysicalTableMapRelationalTable `json:"relationalTables,omitempty" yaml:"relationalTables,omitempty"`

	//
	S3Sources []Quicksight_getDataSetPhysicalTableMapS3Source `json:"s3Sources,omitempty" yaml:"s3Sources,omitempty"`
}
