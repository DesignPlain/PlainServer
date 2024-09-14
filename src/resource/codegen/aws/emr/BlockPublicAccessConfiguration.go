package emr

import types "libds/aws/types"

type BlockPublicAccessConfiguration struct {
	/*
	   Enable or disable EMR Block Public Access.

	   The following arguments are optional:
	*/
	BlockPublicSecurityGroupRules bool `json:"blockPublicSecurityGroupRules,omitempty" yaml:"blockPublicSecurityGroupRules,omitempty"`

	// Configuration block for defining permitted public security group rule port ranges. Can be defined multiple times per resource. Only valid if `block_public_security_group_rules` is set to `true`.
	PermittedPublicSecurityGroupRuleRanges []types.Emr_BlockPublicAccessConfigurationPermittedPublicSecurityGroupRuleRange `json:"permittedPublicSecurityGroupRuleRanges,omitempty" yaml:"permittedPublicSecurityGroupRuleRanges,omitempty"`
}
