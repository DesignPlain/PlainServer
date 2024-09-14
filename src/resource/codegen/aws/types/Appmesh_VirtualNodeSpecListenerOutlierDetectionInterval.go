package types

type Appmesh_VirtualNodeSpecListenerOutlierDetectionInterval struct {
	// Number of time units. Minimum value of `0`.
	Value int `json:"value,omitempty" yaml:"value,omitempty"`

	// Unit of time. Valid values: `ms`, `s`.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`
}
