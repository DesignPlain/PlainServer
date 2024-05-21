package types

type Gkehub_FeatureState struct {
	/*
	   (Output)
	   Output only. The "running state" of the Feature in this Hub.
	   Structure is documented below.
	*/
	States []Gkehub_FeatureStateState `json:"states,omitempty" yaml:"states,omitempty"`
}
