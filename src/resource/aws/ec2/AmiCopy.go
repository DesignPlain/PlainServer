package ec2

import types "DesignSphere_Server/src/resource/aws/types"

type AmiCopy struct {
	// Whether the destination snapshots of the copied image should be encrypted. Defaults to `false`
	Encrypted bool `json:"encrypted,omitempty" yaml:"encrypted,omitempty"`

	// Full ARN of the KMS Key to use when encrypting the snapshots of an image during a copy operation. If not specified, then the default AWS KMS Key will be used
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Region-unique name for the AMI.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Id of the AMI to copy. This id must be valid in the region
	   given by `source_ami_region`.
	*/
	SourceAmiId string `json:"sourceAmiId,omitempty" yaml:"sourceAmiId,omitempty"`

	/*
	   Region from which the AMI will be copied. This may be the
	   same as the AWS provider region in order to create a copy within the same region.
	*/
	SourceAmiRegion string `json:"sourceAmiRegion,omitempty" yaml:"sourceAmiRegion,omitempty"`

	// Date and time to deprecate the AMI. If you specified a value for seconds, Amazon EC2 rounds the seconds to the nearest minute. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	DeprecationTime string `json:"deprecationTime,omitempty" yaml:"deprecationTime,omitempty"`

	// Longer, human-readable description for the AMI.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   ARN of the Outpost to which to copy the AMI.
	   Only specify this parameter when copying an AMI from an AWS Region to an Outpost. The AMI must be in the Region of the destination Outpost.
	*/
	DestinationOutpostArn string `json:"destinationOutpostArn,omitempty" yaml:"destinationOutpostArn,omitempty"`

	/*
	   Nested block describing an EBS block device that should be
	   attached to created instances. The structure of this block is described below.
	*/
	EbsBlockDevices []types.Ec2_AmiCopyEbsBlockDevice `json:"ebsBlockDevices,omitempty" yaml:"ebsBlockDevices,omitempty"`

	/*
	   Nested block describing an ephemeral block device that
	   should be attached to created instances. The structure of this block is described below.
	*/
	EphemeralBlockDevices []types.Ec2_AmiCopyEphemeralBlockDevice `json:"ephemeralBlockDevices,omitempty" yaml:"ephemeralBlockDevices,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
