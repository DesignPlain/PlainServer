package types

type Vpclattice_ServiceNetworkServiceAssociationDnsEntry struct {
	// The ID of the hosted zone.
	HostedZoneId string `json:"hostedZoneId,omitempty" yaml:"hostedZoneId,omitempty"`

	// The domain name of the service.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`
}
