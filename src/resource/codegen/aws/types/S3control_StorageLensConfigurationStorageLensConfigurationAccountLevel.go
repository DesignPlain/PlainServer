package types

type S3control_StorageLensConfigurationStorageLensConfigurationAccountLevel struct {
	// S3 Storage Lens activity metrics. See Activity Metrics below for more details.
	ActivityMetrics S3control_StorageLensConfigurationStorageLensConfigurationAccountLevelActivityMetrics `json:"activityMetrics,omitempty" yaml:"activityMetrics,omitempty"`

	// Advanced cost-optimization metrics for S3 Storage Lens. See Advanced Cost-Optimization Metrics below for more details.
	AdvancedCostOptimizationMetrics S3control_StorageLensConfigurationStorageLensConfigurationAccountLevelAdvancedCostOptimizationMetrics `json:"advancedCostOptimizationMetrics,omitempty" yaml:"advancedCostOptimizationMetrics,omitempty"`

	// Advanced data-protection metrics for S3 Storage Lens. See Advanced Data-Protection Metrics below for more details.
	AdvancedDataProtectionMetrics S3control_StorageLensConfigurationStorageLensConfigurationAccountLevelAdvancedDataProtectionMetrics `json:"advancedDataProtectionMetrics,omitempty" yaml:"advancedDataProtectionMetrics,omitempty"`

	// S3 Storage Lens bucket-level configuration. See Bucket Level below for more details.
	BucketLevel S3control_StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevel `json:"bucketLevel,omitempty" yaml:"bucketLevel,omitempty"`

	// Detailed status code metrics for S3 Storage Lens. See Detailed Status Code Metrics below for more details.
	DetailedStatusCodeMetrics S3control_StorageLensConfigurationStorageLensConfigurationAccountLevelDetailedStatusCodeMetrics `json:"detailedStatusCodeMetrics,omitempty" yaml:"detailedStatusCodeMetrics,omitempty"`
}
