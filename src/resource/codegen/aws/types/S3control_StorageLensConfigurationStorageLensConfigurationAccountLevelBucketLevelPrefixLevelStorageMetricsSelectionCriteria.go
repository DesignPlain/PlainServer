package types

type S3control_StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelPrefixLevelStorageMetricsSelectionCriteria struct {
	// The delimiter of the selection criteria being used.
	Delimiter string `json:"delimiter,omitempty" yaml:"delimiter,omitempty"`

	// The max depth of the selection criteria.
	MaxDepth int `json:"maxDepth,omitempty" yaml:"maxDepth,omitempty"`

	// The minimum number of storage bytes percentage whose metrics will be selected.
	MinStorageBytesPercentage float64 `json:"minStorageBytesPercentage,omitempty" yaml:"minStorageBytesPercentage,omitempty"`
}
