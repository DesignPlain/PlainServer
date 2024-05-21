package types

type Compute_OrganizationSecurityPolicyRuleMatchConfig struct {
	/*
	   Destination IP address range in CIDR format. Required for
	   EGRESS rules.
	*/
	DestIpRanges []string `json:"destIpRanges,omitempty" yaml:"destIpRanges,omitempty"`

	/*
	   Pairs of IP protocols and ports that the rule should match.
	   Structure is documented below.


	   <a name="nested_layer4_config"></a>The `layer4_config` block supports:
	*/
	Layer4Configs []Compute_OrganizationSecurityPolicyRuleMatchConfigLayer4Config `json:"layer4Configs,omitempty" yaml:"layer4Configs,omitempty"`

	/*
	   Source IP address range in CIDR format. Required for
	   INGRESS rules.
	*/
	SrcIpRanges []string `json:"srcIpRanges,omitempty" yaml:"srcIpRanges,omitempty"`
}
