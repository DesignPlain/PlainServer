package s3

import types "libds/aws/types"

type AnalyticsConfiguration struct {
	// Name of the bucket this analytics configuration is associated with.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter types.S3_AnalyticsConfigurationFilter `json:"filter,omitempty" yaml:"filter,omitempty"`

	// Unique identifier of the analytics configuration for the bucket.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Configuration for the analytics data export (documented below).
	StorageClassAnalysis types.S3_AnalyticsConfigurationStorageClassAnalysis `json:"storageClassAnalysis,omitempty" yaml:"storageClassAnalysis,omitempty"`
}
