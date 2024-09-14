package types

type Integrationconnectors_ConnectionEventingConfigRegistrationDestinationConfig struct {
	/*
	   destinations for the connection
	   Structure is documented below.
	*/
	Destinations []Integrationconnectors_ConnectionEventingConfigRegistrationDestinationConfigDestination `json:"destinations,omitempty" yaml:"destinations,omitempty"`

	// Key for the connection
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
}
