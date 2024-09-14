package ec2

import types "libds/aws/types"

type Ami struct {
	// If the image is configured for NitroTPM support, the value is `v2.0`. For more information, see [NitroTPM](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/nitrotpm.html) in the Amazon Elastic Compute Cloud User Guide.
	TpmSupport string `json:"tpmSupport,omitempty" yaml:"tpmSupport,omitempty"`

	// Longer, human-readable description for the AMI.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// If EC2 instances started from this image should require the use of the Instance Metadata Service V2 (IMDSv2), set this argument to `v2.0`. For more information, see [Configure instance metadata options for new instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-IMDS-new-instances.html#configure-IMDS-new-instances-ami-configuration).
	ImdsSupport string `json:"imdsSupport,omitempty" yaml:"imdsSupport,omitempty"`

	// Region-unique name for the AMI.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Keyword to choose what virtualization mode created instances
	   will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	   changes the set of further arguments that are required, as described below.
	*/
	VirtualizationType string `json:"virtualizationType,omitempty" yaml:"virtualizationType,omitempty"`

	//
	SriovNetSupport string `json:"sriovNetSupport,omitempty" yaml:"sriovNetSupport,omitempty"`

	/*
	   Nested block describing an EBS block device that should be
	   attached to created instances. The structure of this block is described below.
	*/
	EbsBlockDevices []types.Ec2_AmiEbsBlockDevice `json:"ebsBlockDevices,omitempty" yaml:"ebsBlockDevices,omitempty"`

	// Whether enhanced networking with ENA is enabled. Defaults to `false`.
	EnaSupport bool `json:"enaSupport,omitempty" yaml:"enaSupport,omitempty"`

	//
	ImageLocation string `json:"imageLocation,omitempty" yaml:"imageLocation,omitempty"`

	// Name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
	RootDeviceName string `json:"rootDeviceName,omitempty" yaml:"rootDeviceName,omitempty"`

	//
	KernelId string `json:"kernelId,omitempty" yaml:"kernelId,omitempty"`

	//
	RamdiskId string `json:"ramdiskId,omitempty" yaml:"ramdiskId,omitempty"`

	// Machine architecture for created instances. Defaults to "x86_64".
	Architecture string `json:"architecture,omitempty" yaml:"architecture,omitempty"`

	// Boot mode of the AMI. For more information, see [Boot modes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-boot.html) in the Amazon Elastic Compute Cloud User Guide.
	BootMode string `json:"bootMode,omitempty" yaml:"bootMode,omitempty"`

	// Date and time to deprecate the AMI. If you specified a value for seconds, Amazon EC2 rounds the seconds to the nearest minute. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	DeprecationTime string `json:"deprecationTime,omitempty" yaml:"deprecationTime,omitempty"`

	/*
	   Nested block describing an ephemeral block device that
	   should be attached to created instances. The structure of this block is described below.
	*/
	EphemeralBlockDevices []types.Ec2_AmiEphemeralBlockDevice `json:"ephemeralBlockDevices,omitempty" yaml:"ephemeralBlockDevices,omitempty"`
}
