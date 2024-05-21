package types

type Compute_PacketMirroringFilter struct {
	/*
	   IP CIDR ranges that apply as a filter on the source (ingress) or
	   destination (egress) IP in the IP header. Only IPv4 is supported.
	*/
	CidrRanges []string `json:"cidrRanges,omitempty" yaml:"cidrRanges,omitempty"`

	/*
	   Direction of traffic to mirror.
	   Default value is `BOTH`.
	   Possible values are: `INGRESS`, `EGRESS`, `BOTH`.
	*/
	Direction string `json:"direction,omitempty" yaml:"direction,omitempty"`

	// Possible IP protocols including tcp, udp, icmp and esp
	IpProtocols []string `json:"ipProtocols,omitempty" yaml:"ipProtocols,omitempty"`
}
