package types

type Networkconnectivity_PolicyBasedRouteFilter struct {
	/*
	   The destination IP range of outgoing packets that this policy-based route applies to. Default is "0.0.0.0/0" if protocol version is IPv4.

	   - - -
	*/
	DestRange string `json:"destRange,omitempty" yaml:"destRange,omitempty"`

	// The IP protocol that this policy-based route applies to. Valid values are 'TCP', 'UDP', and 'ALL'. Default is 'ALL'.
	IpProtocol string `json:"ipProtocol,omitempty" yaml:"ipProtocol,omitempty"`

	/*
	   Internet protocol versions this policy-based route applies to.
	   Possible values are: `IPV4`.
	*/
	ProtocolVersion string `json:"protocolVersion,omitempty" yaml:"protocolVersion,omitempty"`

	// The source IP range of outgoing packets that this policy-based route applies to. Default is "0.0.0.0/0" if protocol version is IPv4.
	SrcRange string `json:"srcRange,omitempty" yaml:"srcRange,omitempty"`
}
