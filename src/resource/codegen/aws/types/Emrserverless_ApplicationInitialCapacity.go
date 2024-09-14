package types

type Emrserverless_ApplicationInitialCapacity struct {
	// The worker type for an analytics framework. For Spark applications, the key can either be set to `Driver` or `Executor`. For Hive applications, it can be set to `HiveDriver` or `TezTask`.
	InitialCapacityType string `json:"initialCapacityType,omitempty" yaml:"initialCapacityType,omitempty"`

	// The initial capacity configuration per worker.
	InitialCapacityConfig Emrserverless_ApplicationInitialCapacityInitialCapacityConfig `json:"initialCapacityConfig,omitempty" yaml:"initialCapacityConfig,omitempty"`
}
