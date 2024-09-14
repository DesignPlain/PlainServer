package types

type Ec2_SpotInstanceRequestEbsBlockDevice struct {
	//
	VolumeId string `json:"volumeId,omitempty" yaml:"volumeId,omitempty"`

	// Whether the volume should be destroyed on instance termination. Defaults to `true`.
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" yaml:"deleteOnTermination,omitempty"`

	// Name of the device to mount.
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	// Enables [EBS encryption](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) on the volume. Defaults to `false`. Cannot be used with `snapshot_id`. Must be configured to perform drift detection.
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	// Amount of provisioned [IOPS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html). Only valid for volume_type of `io1`, `io2` or `gp3`.
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	// Snapshot ID to mount.
	SnapshotId string `json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`

	// Map of tags to assign to the device.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Amazon Resource Name (ARN) of the KMS Key to use when encrypting the volume. Must be configured to perform drift detection.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
	TagsAll map[string]string `json:"tagsAll,omitempty" yaml:"tagsAll,omitempty"`

	// Throughput to provision for a volume in mebibytes per second (MiB/s). This is only valid for `volume_type` of `gp3`.
	Throughput int `json:"throughput,omitempty" yaml:"throughput,omitempty"`

	// Size of the volume in gibibytes (GiB).
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`

	/*
	   Type of volume. Valid values include `standard`, `gp2`, `gp3`, `io1`, `io2`, `sc1`, or `st1`. Defaults to `gp2`.

	   > --NOTE:-- Currently, changes to the `ebs_block_device` configuration of _existing_ resources cannot be automatically detected by this provider. To manage changes and attachments of an EBS block to an instance, use the `aws.ebs.Volume` and `aws.ec2.VolumeAttachment` resources instead. If you use `ebs_block_device` on an `aws.ec2.Instance`, this provider will assume management over the full set of non-root EBS block devices for the instance, treating additional block devices as drift. For this reason, `ebs_block_device` cannot be mixed with external `aws.ebs.Volume` and `aws.ec2.VolumeAttachment` resources for a given instance.
	*/
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`
}
