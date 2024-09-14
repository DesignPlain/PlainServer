package types

type Ec2_AmiCopyEbsBlockDevice struct {
	/*
	   Boolean controlling whether the EBS volumes created to
	   support each created instance will be deleted once that instance is terminated.
	*/
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" yaml:"deleteOnTermination,omitempty"`

	// Boolean controlling whether the created EBS volumes will be encrypted. Can't be used with `snapshot_id`.
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	/*
	   Number of I/O operations per second the
	   created volumes will support.
	*/
	Iops int `json:"iops,omitempty" yaml:"iops,omitempty"`

	/*
	   ID of an EBS snapshot that will be used to initialize the created
	   EBS volumes. If set, the `volume_size` attribute must be at least as large as the referenced
	   snapshot.
	*/
	SnapshotId string `json:"snapshotId,omitempty" yaml:"snapshotId,omitempty"`

	/*
	   Size of created volumes in GiB.
	   If `snapshot_id` is set and `volume_size` is omitted then the volume will have the same size
	   as the selected snapshot.
	*/
	VolumeSize int `json:"volumeSize,omitempty" yaml:"volumeSize,omitempty"`

	// Type of EBS volume to create. Can be `standard`, `gp2`, `gp3`, `io1`, `io2`, `sc1` or `st1` (Default: `standard`).
	VolumeType string `json:"volumeType,omitempty" yaml:"volumeType,omitempty"`

	// Path at which the device is exposed to created instances.
	DeviceName string `json:"deviceName,omitempty" yaml:"deviceName,omitempty"`

	/*
	   ARN of the Outpost on which the snapshot is stored.

	   > --Note:-- You can specify `encrypted` or `snapshot_id` but not both.
	*/
	OutpostArn string `json:"outpostArn,omitempty" yaml:"outpostArn,omitempty"`

	// Throughput that the EBS volume supports, in MiB/s. Only valid for `volume_type` of `gp3`.
	Throughput int `json:"throughput,omitempty" yaml:"throughput,omitempty"`
}
