package types

type Storage_getBucketRetentionPolicy struct {
	// If set to true, the bucket will be locked and permanently restrict edits to the bucket's retention policy.  Caution: Locking a bucket is an irreversible action.
	IsLocked bool `json:"isLocked,omitempty" yaml:"isLocked,omitempty"`

	// The period of time, in seconds, that objects in the bucket must be retained and cannot be deleted, overwritten, or archived. The value must be less than 3,155,760,000 seconds.
	RetentionPeriod int `json:"retentionPeriod,omitempty" yaml:"retentionPeriod,omitempty"`
}
