package types

type Timestreamwrite_TableRetentionProperties struct {
	// The duration for which data must be stored in the memory store. Minimum value of 1. Maximum value of 8766.
	MemoryStoreRetentionPeriodInHours int `json:"memoryStoreRetentionPeriodInHours,omitempty" yaml:"memoryStoreRetentionPeriodInHours,omitempty"`

	// The duration for which data must be stored in the magnetic store. Minimum value of 1. Maximum value of 73000.
	MagneticStoreRetentionPeriodInDays int `json:"magneticStoreRetentionPeriodInDays,omitempty" yaml:"magneticStoreRetentionPeriodInDays,omitempty"`
}
