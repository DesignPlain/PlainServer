package types

type Vpclattice_ListenerDefaultAction struct {
	//
	FixedResponse Vpclattice_ListenerDefaultActionFixedResponse `json:"fixedResponse,omitempty" yaml:"fixedResponse,omitempty"`

	/*
	   Route requests to one or more target groups. See Forward blocks below.

	   > --NOTE:-- You must specify exactly one of the following argument blocks: `fixed_response` or `forward`.
	*/
	Forwards []Vpclattice_ListenerDefaultActionForward `json:"forwards,omitempty" yaml:"forwards,omitempty"`
}
