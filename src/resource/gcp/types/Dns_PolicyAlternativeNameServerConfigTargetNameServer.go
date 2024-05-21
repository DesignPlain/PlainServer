package types

type Dns_PolicyAlternativeNameServerConfigTargetNameServer struct {
	/*
	   Forwarding path for this TargetNameServer. If unset or `default` Cloud DNS will make forwarding
	   decision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go
	   to the Internet. When set to `private`, Cloud DNS will always send queries through VPC for this target
	   Possible values are: `default`, `private`.
	*/
	ForwardingPath string `json:"forwardingPath,omitempty" yaml:"forwardingPath,omitempty"`

	// IPv4 address to forward to.
	Ipv4Address string `json:"ipv4Address,omitempty" yaml:"ipv4Address,omitempty"`
}
