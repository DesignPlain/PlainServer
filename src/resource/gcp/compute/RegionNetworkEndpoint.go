package compute

type RegionNetworkEndpoint struct {
	// Port number of network endpoint.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Region where the containing network endpoint group is located.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   The network endpoint group this endpoint is part of.


	   - - -
	*/
	RegionNetworkEndpointGroup string `json:"regionNetworkEndpointGroup,omitempty" yaml:"regionNetworkEndpointGroup,omitempty"`

	/*
	   Fully qualified domain name of network endpoint.
	   This can only be specified when network_endpoint_type of the NEG is INTERNET_FQDN_PORT.
	*/
	Fqdn string `json:"fqdn,omitempty" yaml:"fqdn,omitempty"`

	/*
	   IPv4 address external endpoint.
	   This can only be specified when network_endpoint_type of the NEG is INTERNET_IP_PORT.
	*/
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`
}
