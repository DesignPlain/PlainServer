package types

type Container_getClusterBinaryAuthorization struct {
	// Enable Binary Authorization for this cluster.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Mode of operation for Binary Authorization policy evaluation.
	EvaluationMode string `json:"evaluationMode,omitempty" yaml:"evaluationMode,omitempty"`
}
