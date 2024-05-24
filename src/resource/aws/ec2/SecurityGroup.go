package ec2

import types "DesignSphere_Server/src/resource/aws/types"

type SecurityGroup struct {
	// Name of the security group. If omitted, the provider will assign a random, unique name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// Instruct the provider to revoke all of the Security Groups attached ingress and egress rules before deleting the rule itself. This is normally not needed, however certain AWS services such as Elastic Map Reduce may automatically add required rules to security groups used with the service, and those rules may contain a cyclic dependency that prevent the security groups from being destroyed without removing the dependency first. Default `false`.
	RevokeRulesOnDelete bool `json:"revokeRulesOnDelete,omitempty" yaml:"revokeRulesOnDelete,omitempty"`

	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// VPC ID. Defaults to the region's default VPC.
	VpcId string `json:"vpcId,omitempty" yaml:"vpcId,omitempty"`

	// Security group description. Defaults to `Managed by Pulumi`. Cannot be `""`. --NOTE--: This field maps to the AWS `GroupDescription` attribute, for which there is no Update API. If you'd like to classify your security groups in a way that can be updated, use `tags`.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Configuration block for egress rules. Can be specified multiple times for each egress rule. Each egress block supports fields documented below. This argument is processed in attribute-as-blocks mode.
	Egress []types.Ec2_SecurityGroupEgress `json:"egress,omitempty" yaml:"egress,omitempty"`

	// Configuration block for ingress rules. Can be specified multiple times for each ingress rule. Each ingress block supports fields documented below. This argument is processed in attribute-as-blocks mode.
	Ingress []types.Ec2_SecurityGroupIngress `json:"ingress,omitempty" yaml:"ingress,omitempty"`
}
