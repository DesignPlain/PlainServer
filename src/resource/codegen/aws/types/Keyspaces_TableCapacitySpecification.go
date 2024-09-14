package types

type Keyspaces_TableCapacitySpecification struct {
	// The throughput capacity specified for read operations defined in read capacity units (RCUs).
	ReadCapacityUnits int `json:"readCapacityUnits,omitempty" yaml:"readCapacityUnits,omitempty"`

	// The read/write throughput capacity mode for a table. Valid values: `PAY_PER_REQUEST`, `PROVISIONED`. The default value is `PAY_PER_REQUEST`.
	ThroughputMode string `json:"throughputMode,omitempty" yaml:"throughputMode,omitempty"`

	// The throughput capacity specified for write operations defined in write capacity units (WCUs).
	WriteCapacityUnits int `json:"writeCapacityUnits,omitempty" yaml:"writeCapacityUnits,omitempty"`
}
