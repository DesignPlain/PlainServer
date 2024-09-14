package types

type Integrationconnectors_ConnectionEventingRuntimeDataStatus struct {
	// An arbitrary description for the Conection.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   (Output)
	   State of the Eventing
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
