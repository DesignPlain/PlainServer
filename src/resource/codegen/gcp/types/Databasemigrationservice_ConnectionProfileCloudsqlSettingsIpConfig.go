package types

type Databasemigrationservice_ConnectionProfileCloudsqlSettingsIpConfig struct {
	/*
	   The list of external networks that are allowed to connect to the instance using the IP.
	   Structure is documented below.
	*/
	AuthorizedNetworks []Databasemigrationservice_ConnectionProfileCloudsqlSettingsIpConfigAuthorizedNetwork `json:"authorizedNetworks,omitempty" yaml:"authorizedNetworks,omitempty"`

	// Whether the instance should be assigned an IPv4 address or not.
	EnableIpv4 bool `json:"enableIpv4,omitempty" yaml:"enableIpv4,omitempty"`

	/*
	   The resource link for the VPC network from which the Cloud SQL instance is accessible for private IP. For example, projects/myProject/global/networks/default.
	   This setting can be updated, but it cannot be removed after it is set.
	*/
	PrivateNetwork string `json:"privateNetwork,omitempty" yaml:"privateNetwork,omitempty"`

	// Whether SSL connections over IP should be enforced or not.
	RequireSsl bool `json:"requireSsl,omitempty" yaml:"requireSsl,omitempty"`
}
