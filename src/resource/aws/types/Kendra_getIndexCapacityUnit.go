package types

type Kendra_getIndexCapacityUnit struct {
	// The amount of extra query capacity for an index and GetQuerySuggestions capacity. For more information, refer to [QueryCapacityUnits](https://docs.aws.amazon.com/kendra/latest/APIReference/API_CapacityUnitsConfiguration.html#Kendra-Type-CapacityUnitsConfiguration-QueryCapacityUnits).
	QueryCapacityUnits int `json:"queryCapacityUnits,omitempty" yaml:"queryCapacityUnits,omitempty"`

	// The amount of extra storage capacity for an index. A single capacity unit provides 30 GB of storage space or 100,000 documents, whichever is reached first. Minimum value of 0.
	StorageCapacityUnits int `json:"storageCapacityUnits,omitempty" yaml:"storageCapacityUnits,omitempty"`
}
