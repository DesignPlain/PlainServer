package types

type Diagflow_CxPageForm struct {
	/*
	   Parameters to collect from the user.
	   Structure is documented below.
	*/
	Parameters []Diagflow_CxPageFormParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}
