package types

type Tpu_NodeSchedulingConfig struct {
	// Defines whether the TPU instance is preemptible.
	Preemptible bool `json:"preemptible,omitempty" yaml:"preemptible,omitempty"`
}
