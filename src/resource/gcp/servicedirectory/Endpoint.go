package servicedirectory

type Endpoint struct {
	// IPv4 or IPv6 address of the endpoint.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	/*
	   The Resource ID must be 1-63 characters long, including digits,
	   lowercase letters or the hyphen character.


	   - - -
	*/
	EndpointId string `json:"endpointId,omitempty" yaml:"endpointId,omitempty"`

	/*
	   Metadata for the endpoint. This data can be consumed
	   by service clients. The entire metadata dictionary may contain
	   up to 512 characters, spread across all key-value pairs.
	   Metadata that goes beyond any these limits will be rejected.
	*/
	Metadata map[string]string `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// The URL to the network, such as projects/PROJECT_NUMBER/locations/global/networks/NETWORK_NAME.
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   Port that the endpoint is running on, must be in the
	   range of [0, 65535]. If unspecified, the default is 0.
	*/
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// The resource name of the service that this endpoint provides.
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}
