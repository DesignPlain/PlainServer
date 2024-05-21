package types

type Redis_ClusterDiscoveryEndpoint struct {
	// Output only. The IP allocated on the consumer network for the PSC forwarding rule.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	// Output only. The port number of the exposed Redis endpoint.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	/*
	   Output only. Customer configuration for where the endpoint
	   is created and accessed from.
	   Structure is documented below.
	*/
	PscConfig Redis_ClusterDiscoveryEndpointPscConfig `json:"pscConfig,omitempty" yaml:"pscConfig,omitempty"`
}
