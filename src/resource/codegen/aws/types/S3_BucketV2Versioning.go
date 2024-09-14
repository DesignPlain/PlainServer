package types

type S3_BucketV2Versioning struct {
	// Enable versioning. Once you version-enable a bucket, it can never return to an unversioned state. You can, however, suspend versioning on that bucket.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Enable MFA delete for either `Change the versioning state of your bucket` or `Permanently delete an object version`. Default is `false`. This cannot be used to toggle this setting but is available to allow managed buckets to reflect the state in AWS
	MfaDelete bool `json:"mfaDelete,omitempty" yaml:"mfaDelete,omitempty"`
}
