package route53

import types "DesignSphere_Server/src/resource/aws/types"

type Zone struct {
	// A comment for the hosted zone. Defaults to 'Managed by Pulumi'.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
	DelegationSetId string `json:"delegationSetId,omitempty" yaml:"delegationSetId,omitempty"`

	// Whether to destroy all records (possibly managed outside of this provider) in the zone when destroying the zone.
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	// This is the name of the hosted zone.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A mapping of tags to assign to the zone. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegation_set_id` argument in this resource and any `aws.route53.ZoneAssociation` resource specifying the same zone ID. Detailed below.
	Vpcs []types.Route53_ZoneVpc `json:"vpcs,omitempty" yaml:"vpcs,omitempty"`
}
