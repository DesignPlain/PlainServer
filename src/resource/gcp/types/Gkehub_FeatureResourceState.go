package types

type Gkehub_FeatureResourceState struct {
	/*
	   (Output)
	   Output only. The "running state" of the Feature in this Hub.
	   Structure is documented below.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	/*
	   (Output)
	   Whether this Feature has outstanding resources that need to be cleaned up before it can be disabled.
	*/
	HasResources bool `json:"hasResources,omitempty" yaml:"hasResources,omitempty"`
}
