package types

type Timestreamwrite_getTableSchemaCompositePartitionKey struct {
	//
	EnforcementInRecord string `json:"enforcementInRecord,omitempty" yaml:"enforcementInRecord,omitempty"`

	// Name of the Timestream table.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Type of partition key.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
