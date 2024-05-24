package ec2

import types "DesignSphere_Server/src/resource/aws/types"

type DefaultNetworkAcl struct {
	// Configuration block for an ingress rule. Detailed below.
	Ingress []types.Ec2_DefaultNetworkAclIngress `json:"ingress,omitempty" yaml:"ingress,omitempty"`

	// List of Subnet IDs to apply the ACL to. See the notes above on Managing Subnets in the Default Network ACL
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Network ACL ID to manage. This attribute is exported from `aws.ec2.Vpc`, or manually found via the AWS Console.

	   The following arguments are optional:
	*/
	DefaultNetworkAclId string `json:"defaultNetworkAclId,omitempty" yaml:"defaultNetworkAclId,omitempty"`

	// Configuration block for an egress rule. Detailed below.
	Egress []types.Ec2_DefaultNetworkAclEgress `json:"egress,omitempty" yaml:"egress,omitempty"`
}
