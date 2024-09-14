package types

type Integrationconnectors_ConnectionDestinationConfig struct {
	/*
	   The destinations for the key.
	   Structure is documented below.
	*/
	Destinations []Integrationconnectors_ConnectionDestinationConfigDestination `json:"destinations,omitempty" yaml:"destinations,omitempty"`

	// The key is the destination identifier that is supported by the Connector.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
