package types

type Ec2_SpotInstanceRequestRootBlockDevice struct {
	/*
	   Type of volume. Valid values include `standard`, `gp2`, `gp3`, `io1`, `io2`, `sc1`, or `st1`. Defaults to the volume type that the AMI uses.

	   Modifying the `encrypted` or `kms_key_id` settings of the `root_block_device` requires resource replacement.
	*/
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`

	// Whether the volume should be destroyed on instance termination. Defaults to `true`.
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" yaml:"deleteOnTermination,omitempty"`

	// Whether to enable volume encryption. Defaults to `false`. Must be configured to perform drift detection.
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	// Amount of provisioned [IOPS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html). Only valid for volume_type of `io1`, `io2` or `gp3`.
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	// Amazon Resource Name (ARN) of the KMS Key to use when encrypting the volume. Must be configured to perform drift detection.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Map of tags to assign to the device.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
	TagsAll map[string]string `json:"tagsAll,omitempty" yaml:"tagsAll,omitempty"`

	// Throughput to provision for a volume in mebibytes per second (MiB/s). This is only valid for `volume_type` of `gp3`.
	Throughput int `json:"throughput,omitempty" yaml:"throughput,omitempty"`

	// Name of the device to mount.
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	//
	VolumeId string `json:"volumeId,omitempty" yaml:"volumeId,omitempty"`

	// Size of the volume in gibibytes (GiB).
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`
}
