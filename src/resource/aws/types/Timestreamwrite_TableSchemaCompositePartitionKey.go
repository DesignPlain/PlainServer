package types

type Timestreamwrite_TableSchemaCompositePartitionKey struct {
	// The level of enforcement for the specification of a dimension key in ingested records. Valid values: `REQUIRED`, `OPTIONAL`.
	EnforcementInRecord string `json:"enforcementInRecord,omitempty" yaml:"enforcementInRecord,omitempty"`

	// The name of the attribute used for a dimension key.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The type of the partition key. Valid values: `DIMENSION`, `MEASURE`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
