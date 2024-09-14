package types

type S3control_BucketLifecycleConfigurationRuleExpiration struct {
	// Date the object is to be deleted. Should be in `YYYY-MM-DD` date format, e.g., `2020-09-30`.
	Date string `json:"date,omitempty" yaml:"date,omitempty"`

	// Number of days before the object is to be deleted.
	Days int `json:"days,omitempty" yaml:"days,omitempty"`

	// Enable to remove a delete marker with no noncurrent versions. Cannot be specified with `date` or `days`.
	ExpiredObjectDeleteMarker bool `json:"expiredObjectDeleteMarker,omitempty" yaml:"expiredObjectDeleteMarker,omitempty"`
}
