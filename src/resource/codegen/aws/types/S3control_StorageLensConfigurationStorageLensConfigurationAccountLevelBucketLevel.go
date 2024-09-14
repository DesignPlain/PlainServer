package types

type S3control_StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevel struct {
	// S3 Storage Lens activity metrics. See Activity Metrics above for more details.
	ActivityMetrics S3control_StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelActivityMetrics `json:"activityMetrics,omitempty" yaml:"activityMetrics,omitempty"`

	// Advanced cost-optimization metrics for S3 Storage Lens. See Advanced Cost-Optimization Metrics above for more details.
	AdvancedCostOptimizationMetrics S3control_StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetrics `json:"advancedCostOptimizationMetrics,omitempty" yaml:"advancedCostOptimizationMetrics,omitempty"`

	// Advanced data-protection metrics for S3 Storage Lens. See Advanced Data-Protection Metrics above for more details.
	AdvancedDataProtectionMetrics S3control_StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetrics `json:"advancedDataProtectionMetrics,omitempty" yaml:"advancedDataProtectionMetrics,omitempty"`

	// Detailed status code metrics for S3 Storage Lens. See Detailed Status Code Metrics above for more details.
	DetailedStatusCodeMetrics S3control_StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetrics `json:"detailedStatusCodeMetrics,omitempty" yaml:"detailedStatusCodeMetrics,omitempty"`

	// Prefix-level metrics for S3 Storage Lens. See Prefix Level below for more details.
	PrefixLevel S3control_StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelPrefixLevel `json:"prefixLevel,omitempty" yaml:"prefixLevel,omitempty"`
}
