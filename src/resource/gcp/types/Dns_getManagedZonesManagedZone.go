package types

type Dns_getManagedZonesManagedZone struct {
	// The fully qualified DNS name of this zone.
	DnsName string `json:"dnsName,omitempty" yaml:"dnsName,omitempty"`

	// DNS managed zone identifier
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Unique identifier for the resource; defined by the server.
	ManagedZoneId int `json:"managedZoneId,omitempty" yaml:"managedZoneId,omitempty"`

	// A unique name for the resource.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The list of nameservers that will be authoritative for this domain. Use NS records to redirect from your DNS provider to these names, thus making Google Cloud DNS authoritative for this zone.
	NameServers []string `json:"nameServers,omitempty" yaml:"nameServers,omitempty"`

	// The ID of the project containing Google Cloud DNS zones. If this is not provided the default project will be used.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private Cloud resources.
	Visibility string `json:"visibility,omitempty" yaml:"visibility,omitempty"`

	// A textual description field.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
