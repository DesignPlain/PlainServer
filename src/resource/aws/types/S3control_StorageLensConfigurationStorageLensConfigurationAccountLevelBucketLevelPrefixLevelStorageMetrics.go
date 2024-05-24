package types

type S3control_StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetrics struct {
	// Whether prefix-level storage metrics are enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Selection criteria. See Selection Criteria below for more details.
	SelectionCriteria S3control_StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsSelectionCriteria `json:"selectionCriteria,omitempty" yaml:"selectionCriteria,omitempty"`
}
