package activedirectory

type Domain struct {
	/*
	   Resource labels that can contain user-provided metadata
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Locations where domain needs to be provisioned. [regions][compute/docs/regions-zones/]
	   e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	*/
	Locations []string `json:"locations,omitempty" yaml:"locations,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger.
	   Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
	*/
	ReservedIpRange string `json:"reservedIpRange,omitempty" yaml:"reservedIpRange,omitempty"`

	/*
	   The name of delegated administrator account used to perform Active Directory operations.
	   If not specified, setupadmin will be used.
	*/
	Admin string `json:"admin,omitempty" yaml:"admin,omitempty"`

	/*
	   The full names of the Google Compute Engine networks the domain instance is connected to. The domain is only available on networks listed in authorizedNetworks.
	   If CIDR subnets overlap between networks, domain creation will fail.
	*/
	AuthorizedNetworks []string `json:"authorizedNetworks,omitempty" yaml:"authorizedNetworks,omitempty"`

	/*
	   The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,
	   https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.


	   - - -
	*/
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`
}
