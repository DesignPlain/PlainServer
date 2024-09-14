package ebs

type Volume struct {
	// If true, snapshot will be created before volume deletion. Any tags on the volume will be migrated to the snapshot. By default set to false
	FinalSnapshot bool `json:"finalSnapshot,omitempty" yaml:"finalSnapshot,omitempty"`

	// The amount of IOPS to provision for the disk. Only valid for `type` of `io1`, `io2` or `gp3`.
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	// The ARN for the KMS encryption key. When specifying `kms_key_id`, `encrypted` needs to be set to true. Note: The provider must be running with credentials which have the `GenerateDataKeyWithoutPlaintext` permission on the specified KMS key as required by the [EBS KMS CMK volume provisioning process](https://docs.aws.amazon.com/kms/latest/developerguide/services-ebs.html#ebs-cmk) to prevent a volume from being created and almost immediately deleted.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The type of EBS volume. Can be `standard`, `gp2`, `gp3`, `io1`, `io2`, `sc1` or `st1` (Default: `gp2`).
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// If true, the disk will be encrypted.
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	// Specifies whether to enable Amazon EBS Multi-Attach. Multi-Attach is supported on `io1` and `io2` volumes.
	MultiAttachEnabled bool `json:"multiAttachEnabled,omitempty" yaml:"multiAttachEnabled,omitempty"`

	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn string `json:"outpostArn,omitempty" yaml:"outpostArn,omitempty"`

	// The size of the drive in GiBs.
	Size int `json:"size,omitempty" yaml:"size,omitempty"`

	// A snapshot to base the EBS volume off of.
	SnapshotId string `json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`

	/*
	   The throughput that the volume supports, in MiB/s. Only valid for `type` of `gp3`.

	   > --NOTE:-- When changing the `size`, `iops` or `type` of an instance, there are [considerations](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/considerations.html) to be aware of.
	*/
	Throughput int `json:"throughput,omitempty" yaml:"throughput,omitempty"`

	// The AZ where the EBS volume will exist.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`
}
