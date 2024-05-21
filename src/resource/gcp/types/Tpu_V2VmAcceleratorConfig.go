package types

type Tpu_V2VmAcceleratorConfig struct {
	// Topology of TPU in chips.
	Topology string `json:"topology,omitempty" yaml:"topology,omitempty"`

	/*
	   Type of TPU.
	   Possible values are: `V2`, `V3`, `V4`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
