package types

type Ec2_LaunchTemplateBlockDeviceMappingEbs struct {
	/*
	   The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use when creating the encrypted volume.
	   `encrypted` must be set to `true` when this is set.
	*/
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// The Snapshot ID to mount.
	SnapshotId string `json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`

	// The throughput to provision for a `gp3` volume in MiB/s (specified as an integer, e.g., 500), with a maximum of 1,000 MiB/s.
	Throughput int `json:"throughput,omitempty" yaml:"throughput,omitempty"`

	// The size of the volume in gigabytes.
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`

	/*
	   The volume type.
	   Can be one of `standard`, `gp2`, `gp3`, `io1`, `io2`, `sc1` or `st1`.
	*/
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`

	/*
	   Whether the volume should be destroyed on instance termination.
	   See [Preserving Amazon EBS Volumes on Instance Termination](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/preserving-volumes-on-termination.html) for more information.
	*/
	DeleteOnTermination string `json:"deleteOnTermination,omitempty" yaml:"deleteOnTermination,omitempty"`

	/*
	   Enables [EBS encryption](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) on the volume.
	   Cannot be used with `snapshot_id`.
	*/
	Encrypted string `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	/*
	   The amount of provisioned [IOPS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html).
	   This must be set with a `volume_type` of `"io1/io2/gp3"`.
	*/
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`
}
