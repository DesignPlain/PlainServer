package types

type Compute_SecurityPolicyRuleMatchConfig struct {
	/*
	   Set of IP addresses or ranges (IPV4 or IPV6) in CIDR notation
	   to match against inbound traffic. There is a limit of 10 IP ranges per rule. A value of `-` matches all IPs
	   (can be used to override the default behavior).
	*/
	SrcIpRanges []string `json:"srcIpRanges,omitempty" yaml:"srcIpRanges,omitempty"`
}
