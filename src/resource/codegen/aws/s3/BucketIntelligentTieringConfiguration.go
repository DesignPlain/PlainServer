package s3

import types "libds/aws/types"

type BucketIntelligentTieringConfiguration struct {
	// Name of the bucket this intelligent tiering configuration is associated with.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Bucket filter. The configuration only includes objects that meet the filter's criteria (documented below).
	Filter types.S3_BucketIntelligentTieringConfigurationFilter `json:"filter,omitempty" yaml:"filter,omitempty"`

	// Unique name used to identify the S3 Intelligent-Tiering configuration for the bucket.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies the status of the configuration. Valid values: `Enabled`, `Disabled`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// S3 Intelligent-Tiering storage class tiers of the configuration (documented below).
	Tierings []types.S3_BucketIntelligentTieringConfigurationTiering `json:"tierings,omitempty" yaml:"tierings,omitempty"`
}
