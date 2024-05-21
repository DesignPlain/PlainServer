package types

type Compute_FirewallPolicyRuleMatch struct {
	// Domain names that will be used to match against the resolved domain name of destination of traffic. Can only be specified if DIRECTION is egress.
	DestFqdns []string `json:"destFqdns,omitempty" yaml:"destFqdns,omitempty"`

	// Pairs of IP protocols and ports that the rule should match.
	Layer4Configs []Compute_FirewallPolicyRuleMatchLayer4Config `json:"layer4Configs,omitempty" yaml:"layer4Configs,omitempty"`

	/*
	   Name of the Google Cloud Threat Intelligence list.

	   The `layer4_configs` block supports:
	*/
	SrcThreatIntelligences []string `json:"srcThreatIntelligences,omitempty" yaml:"srcThreatIntelligences,omitempty"`

	// The Unicode country codes whose IP addresses will be used to match against the source of traffic. Can only be specified if DIRECTION is ingress.
	SrcRegionCodes []string `json:"srcRegionCodes,omitempty" yaml:"srcRegionCodes,omitempty"`

	// Address groups which should be matched against the traffic destination. Maximum number of destination address groups is 10. Destination address groups is only supported in Egress rules.
	DestAddressGroups []string `json:"destAddressGroups,omitempty" yaml:"destAddressGroups,omitempty"`

	// CIDR IP address range. Maximum number of destination CIDR IP ranges allowed is 256.
	DestIpRanges []string `json:"destIpRanges,omitempty" yaml:"destIpRanges,omitempty"`

	// The Unicode country codes whose IP addresses will be used to match against the source of traffic. Can only be specified if DIRECTION is egress.
	DestRegionCodes []string `json:"destRegionCodes,omitempty" yaml:"destRegionCodes,omitempty"`

	// Name of the Google Cloud Threat Intelligence list.
	DestThreatIntelligences []string `json:"destThreatIntelligences,omitempty" yaml:"destThreatIntelligences,omitempty"`

	// Address groups which should be matched against the traffic source. Maximum number of source address groups is 10. Source address groups is only supported in Ingress rules.
	SrcAddressGroups []string `json:"srcAddressGroups,omitempty" yaml:"srcAddressGroups,omitempty"`

	// Domain names that will be used to match against the resolved domain name of source of traffic. Can only be specified if DIRECTION is ingress.
	SrcFqdns []string `json:"srcFqdns,omitempty" yaml:"srcFqdns,omitempty"`

	// CIDR IP address range. Maximum number of source CIDR IP ranges allowed is 256.
	SrcIpRanges []string `json:"srcIpRanges,omitempty" yaml:"srcIpRanges,omitempty"`
}
