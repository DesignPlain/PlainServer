package compute

type GlobalNetworkEndpoint struct {
	// Port number of the external endpoint.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Fully qualified domain name of network endpoint.
	   This can only be specified when network_endpoint_type of the NEG is INTERNET_FQDN_PORT.
	*/
	Fqdn string `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`

	/*
	   The global network endpoint group this endpoint is part of.


	   - - -
	*/
	GlobalNetworkEndpointGroup string `json:"globalNetworkEndpointGroup,omitempty" yaml:"globalNetworkEndpointGroup,omitempty"`

	// IPv4 address external endpoint.
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`
}
