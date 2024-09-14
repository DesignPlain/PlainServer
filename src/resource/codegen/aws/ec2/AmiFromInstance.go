package ec2

import types "libds/aws/types"

type AmiFromInstance struct {
	/*
	   Boolean that overrides the behavior of stopping
	   the instance before snapshotting. This is risky since it may cause a snapshot of an
	   inconsistent filesystem state, but can be used to avoid downtime if the user otherwise
	   guarantees that no filesystem writes will be underway at the time of snapshot.
	*/
	SnapshotWithoutReboot bool `json:"snapshotWithoutReboot,omitempty" yaml:"snapshotWithoutReboot,omitempty"`

	// ID of the instance to use as the basis of the AMI.
	SourceInstanceId string `json:"sourceInstanceId,omitempty" yaml:"sourceInstanceId,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Date and time to deprecate the AMI. If you specified a value for seconds, Amazon EC2 rounds the seconds to the nearest minute. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	DeprecationTime string `json:"deprecationTime,omitempty" yaml:"deprecationTime,omitempty"`

	// Longer, human-readable description for the AMI.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Nested block describing an EBS block device that should be
	   attached to created instances. The structure of this block is described below.
	*/
	EbsBlockDevices []types.Ec2_AmiFromInstanceEbsBlockDevice `json:"ebsBlockDevices,omitempty" yaml:"ebsBlockDevices,omitempty"`

	/*
	   Nested block describing an ephemeral block device that
	   should be attached to created instances. The structure of this block is described below.
	*/
	EphemeralBlockDevices []types.Ec2_AmiFromInstanceEphemeralBlockDevice `json:"ephemeralBlockDevices,omitempty" yaml:"ephemeralBlockDevices,omitempty"`

	// Region-unique name for the AMI.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
