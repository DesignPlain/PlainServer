package types

type Endpoints_ServiceEndpoint struct {
	// The FQDN of the endpoint as described in the config.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	// The simple name of the endpoint as described in the config.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
