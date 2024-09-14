package types

type S3_BucketVersioningV2VersioningConfiguration struct {
	// Specifies whether MFA delete is enabled in the bucket versioning configuration. Valid values: `Enabled` or `Disabled`.
	MfaDelete string `json:"mfaDelete,omitempty" yaml:"mfaDelete,omitempty"`

	// Versioning state of the bucket. Valid values: `Enabled`, `Suspended`, or `Disabled`. `Disabled` should only be used when creating or importing resources that correspond to unversioned S3 buckets.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
