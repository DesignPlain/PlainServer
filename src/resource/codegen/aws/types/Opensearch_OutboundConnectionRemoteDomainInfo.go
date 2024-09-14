package types

type Opensearch_OutboundConnectionRemoteDomainInfo struct {
	// The Account ID of the owner of the remote domain.
	OwnerId string `json:"ownerId,omitempty" yaml:"ownerId,omitempty"`

	// The region of the remote domain.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// The name of the remote domain.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`
}
