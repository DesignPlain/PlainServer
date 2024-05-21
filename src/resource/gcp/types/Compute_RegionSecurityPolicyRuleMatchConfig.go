package types

type Compute_RegionSecurityPolicyRuleMatchConfig struct {
	// CIDR IP address range. Maximum number of srcIpRanges allowed is 10.
	SrcIpRanges []string `json:"srcIpRanges,omitempty" yaml:"srcIpRanges,omitempty"`
}
