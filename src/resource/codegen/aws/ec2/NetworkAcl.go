package ec2

import types "libds/aws/types"

type NetworkAcl struct {
	// Specifies an egress rule. Parameters defined below.
	Egress []types.Ec2_NetworkAclEgress `json:"egress,omitempty" yaml:"egress,omitempty"`

	// Specifies an ingress rule. Parameters defined below.
	Ingress []types.Ec2_NetworkAclIngress `json:"ingress,omitempty" yaml:"ingress,omitempty"`

	// A list of Subnet IDs to apply the ACL to
	SubnetIds []string `json:"subnetIds,omitempty" yaml:"subnetIds,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ID of the associated VPC.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`
}
