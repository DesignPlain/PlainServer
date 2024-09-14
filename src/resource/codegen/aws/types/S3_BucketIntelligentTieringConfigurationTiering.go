package types

type S3_BucketIntelligentTieringConfigurationTiering struct {
	// S3 Intelligent-Tiering access tier. Valid values: `ARCHIVE_ACCESS`, `DEEP_ARCHIVE_ACCESS`.
	AccessTier string `json:"accessTier,omitempty" yaml:"accessTier,omitempty"`

	// Number of consecutive days of no access after which an object will be eligible to be transitioned to the corresponding tier.
	Days int `json:"days,omitempty" yaml:"days,omitempty"`
}
