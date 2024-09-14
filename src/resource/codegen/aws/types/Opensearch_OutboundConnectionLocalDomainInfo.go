package types

type Opensearch_OutboundConnectionLocalDomainInfo struct {
	// The region of the local domain.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The name of the local domain.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// The Account ID of the owner of the local domain.
	OwnerId string `json:"ownerId,omitempty" yaml:"ownerId,omitempty"`
}
