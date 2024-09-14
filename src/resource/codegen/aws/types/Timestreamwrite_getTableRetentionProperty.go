package types

type Timestreamwrite_getTableRetentionProperty struct {
	// Duration in hours in which the data must be stored in memory store.
	MemoryStoreRetentionPeriodInHours int `json:"memoryStoreRetentionPeriodInHours,omitempty" yaml:"memoryStoreRetentionPeriodInHours,omitempty"`

	// Duration in days in which the data must be stored in magnetic store.
	MagneticStoreRetentionPeriodInDays int `json:"magneticStoreRetentionPeriodInDays,omitempty" yaml:"magneticStoreRetentionPeriodInDays,omitempty"`
}
