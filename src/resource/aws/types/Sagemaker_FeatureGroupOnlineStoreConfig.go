package types

type Sagemaker_FeatureGroupOnlineStoreConfig struct {
	// Option for different tiers of low latency storage for real-time data retrieval. Valid values are `Standard`, or `InMemory`.
	StorageType string `json:"storageType,omitempty" yaml:"storageType,omitempty"`

	// Time to live duration, where the record is hard deleted after the expiration time is reached; ExpiresAt = EventTime + TtlDuration.. See TTl Duration Below.
	TtlDuration Sagemaker_FeatureGroupOnlineStoreConfigTtlDuration `json:"ttlDuration,omitempty" yaml:"ttlDuration,omitempty"`

	// Set to `true` to disable the automatic creation of an AWS Glue table when configuring an OfflineStore.
	EnableOnlineStore bool `json:"enableOnlineStore,omitempty" yaml:"enableOnlineStore,omitempty"`

	// Security config for at-rest encryption of your OnlineStore. See Security Config Below.
	SecurityConfig Sagemaker_FeatureGroupOnlineStoreConfigSecurityConfig `json:"securityConfig,omitempty" yaml:"securityConfig,omitempty"`
}
