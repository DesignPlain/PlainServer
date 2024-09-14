package types

type S3_DirectoryBucketLocation struct {
	// Location type. Valid values: `AvailabilityZone`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// [Availability Zone ID](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html#az-ids).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
