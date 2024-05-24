package types

type Vpclattice_getListenerDefaultAction struct {
	//
	FixedResponses []Vpclattice_getListenerDefaultActionFixedResponse `json:"fixedResponses,omitempty" yaml:"fixedResponses,omitempty"`

	//
	Forwards []Vpclattice_getListenerDefaultActionForward `json:"forwards,omitempty" yaml:"forwards,omitempty"`
}
