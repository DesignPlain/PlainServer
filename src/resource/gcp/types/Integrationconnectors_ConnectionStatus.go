package types

type Integrationconnectors_ConnectionStatus struct {
	// An arbitrary description for the Conection.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   (Output)
	   State of the Eventing
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	/*
	   (Output)
	   Current status of eventing.
	   Structure is documented below.
	*/
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
