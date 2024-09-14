package types

type Tpu_V2VmSchedulingConfig struct {
	// Defines whether the node is preemptible.
	Preemptible bool `json:"preemptible,omitempty" yaml:"preemptible,omitempty"`

	// Whether the node is created under a reservation.
	Reserved bool `json:"reserved,omitempty" yaml:"reserved,omitempty"`
}
