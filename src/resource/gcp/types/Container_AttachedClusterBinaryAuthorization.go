package types

type Container_AttachedClusterBinaryAuthorization struct {
	/*
	   Configure Binary Authorization evaluation mode.
	   Possible values are: `DISABLED`, `PROJECT_SINGLETON_POLICY_ENFORCE`.
	*/
	EvaluationMode string `json:"evaluationMode,omitempty" yaml:"evaluationMode,omitempty"`
}
