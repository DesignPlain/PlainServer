package types

type Networkmanager_getCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocation struct {
	//
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// ASN of the Core Network Edge in an AWS Region. By default, the ASN will be a single integer automatically assigned from `asn_ranges`
	Asn string `json:"asn,omitempty" yaml:"asn,omitempty"`

	// The local CIDR blocks for this Core Network Edge for AWS Transit Gateway Connect attachments. By default, this CIDR block will be one or more optional IPv4 and IPv6 CIDR prefixes auto-assigned from `inside_cidr_blocks`.
	InsideCidrBlocks []string `json:"insideCidrBlocks,omitempty" yaml:"insideCidrBlocks,omitempty"`
}
