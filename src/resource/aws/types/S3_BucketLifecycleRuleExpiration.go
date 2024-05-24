package types

type S3_BucketLifecycleRuleExpiration struct {
	// Specifies the date after which you want the corresponding action to take effect.
	Date string `json:"date,omitempty" yaml:"date,omitempty"`

	// Specifies the number of days after object creation when the specific rule action takes effect.
	Days int `json:"days,omitempty" yaml:"days,omitempty"`

	// On a versioned bucket (versioning-enabled or versioning-suspended bucket), you can add this element in the lifecycle configuration to direct Amazon S3 to delete expired object delete markers. This cannot be specified with Days or Date in a Lifecycle Expiration Policy.
	ExpiredObjectDeleteMarker bool `json:"expiredObjectDeleteMarker,omitempty" yaml:"expiredObjectDeleteMarker,omitempty"`
}
