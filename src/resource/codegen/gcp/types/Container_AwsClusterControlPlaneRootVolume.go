package types

type Container_AwsClusterControlPlaneRootVolume struct {
	// Optional. The Amazon Resource Name (ARN) of the Customer Managed Key (CMK) used to encrypt AWS EBS volumes. If not specified, the default Amazon managed key associated to the AWS region where this cluster runs will be used.
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`

	// Optional. The size of the volume, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.
	SizeGib int `json:"sizeGib,omitempty" yaml:"sizeGib,omitempty"`

	// Optional. The throughput to provision for the volume, in MiB/s. Only valid if the volume type is GP3. If volume type is gp3 and throughput is not specified, the throughput will defaults to 125.
	Throughput int `json:"throughput,omitempty" yaml:"throughput,omitempty"`

	// Optional. Type of the EBS volume. When unspecified, it defaults to GP2 volume. Possible values: VOLUME_TYPE_UNSPECIFIED, GP2, GP3
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`

	// Optional. The number of I/O operations per second (IOPS) to provision for GP3 volume.
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`
}
