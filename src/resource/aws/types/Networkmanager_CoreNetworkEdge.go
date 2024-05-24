package types

type Networkmanager_CoreNetworkEdge struct {
	// Inside IP addresses used for core network edges.
	InsideCidrBlocks []string `json:"insideCidrBlocks,omitempty" yaml:"insideCidrBlocks,omitempty"`

	// ASN of a core network edge.
	Asn int `json:"asn,omitempty" yaml:"asn,omitempty"`

	// Region where a core network edge is located.
	EdgeLocation string `json:"edgeLocation,omitempty" yaml:"edgeLocation,omitempty"`
}
