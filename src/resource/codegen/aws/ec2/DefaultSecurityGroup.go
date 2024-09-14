package ec2

import types "libds/aws/types"

type DefaultSecurityGroup struct {
	//
	RevokeRulesOnDelete bool `json:"revokeRulesOnDelete,omitempty" yaml:"revokeRulesOnDelete,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// VPC ID. --Note that changing the `vpc_id` will _not_ restore any default security group rules that were modified, added, or removed.-- It will be left in its current state.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// Configuration block. Detailed below.
	Egress []types.Ec2_DefaultSecurityGroupEgress `json:"egress,omitempty" yaml:"egress,omitempty"`

	// Configuration block. Detailed below.
	Ingress []types.Ec2_DefaultSecurityGroupIngress `json:"ingress,omitempty" yaml:"ingress,omitempty"`
}
