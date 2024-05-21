package activedirectory

type DomainTrust struct {
	// Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
	SelectiveAuthentication bool `json:"selectiveAuthentication,omitempty" yaml:"selectiveAuthentication,omitempty"`

	// The target DNS server IP addresses which can resolve the remote domain involved in the trust.
	TargetDnsIpAddresses []string `json:"targetDnsIpAddresses,omitempty" yaml:"targetDnsIpAddresses,omitempty"`

	// The fully qualified target domain name which will be in trust with the current domain.
	TargetDomainName string `json:"targetDomainName,omitempty" yaml:"targetDomainName,omitempty"`

	/*
	   The trust direction, which decides if the current domain is trusted, trusting, or both.
	   Possible values are: `INBOUND`, `OUTBOUND`, `BIDIRECTIONAL`.
	*/
	TrustDirection string `json:"trustDirection,omitempty" yaml:"trustDirection,omitempty"`

	/*
	   The trust secret used for the handshake with the target domain. This will not be stored.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	TrustHandshakeSecret string `json:"trustHandshakeSecret,omitempty" yaml:"trustHandshakeSecret,omitempty"`

	/*
	   The type of trust represented by the trust resource.
	   Possible values are: `FOREST`, `EXTERNAL`.
	*/
	TrustType string `json:"trustType,omitempty" yaml:"trustType,omitempty"`

	/*
	   The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
	   https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.


	   - - -
	*/
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
