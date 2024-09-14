package types

type Container_ClusterBinaryAuthorization struct {
	/*
	   Mode of operation for Binary Authorization policy evaluation. Valid values are `DISABLED`
	   and `PROJECT_SINGLETON_POLICY_ENFORCE`.
	*/
	EvaluationMode string `json:"evaluationMode,omitempty" yaml:"evaluationMode,omitempty"`

	// Enable Binary Authorization for this cluster. Deprecated in favor of `evaluation_mode`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
