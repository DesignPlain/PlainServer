package types

type S3_BucketLifecycleConfigurationV2RuleExpiration struct {
	// Date the object is to be moved or deleted. The date value must be in [RFC3339 full-date format](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) e.g. `2023-08-22`.
	Date string `json:"date,omitempty" yaml:"date,omitempty"`

	// Lifetime, in days, of the objects that are subject to the rule. The value must be a non-zero positive integer.
	Days int `json:"days,omitempty" yaml:"days,omitempty"`

	// Indicates whether Amazon S3 will remove a delete marker with no noncurrent versions. If set to `true`, the delete marker will be expired; if set to `false` the policy takes no action.
	ExpiredObjectDeleteMarker bool `json:"expiredObjectDeleteMarker,omitempty" yaml:"expiredObjectDeleteMarker,omitempty"`
}
