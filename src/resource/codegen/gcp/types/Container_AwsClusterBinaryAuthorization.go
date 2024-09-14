package types

type Container_AwsClusterBinaryAuthorization struct {
	// Mode of operation for Binary Authorization policy evaluation. Possible values: DISABLED, PROJECT_SINGLETON_POLICY_ENFORCE
	EvaluationMode string `json:"evaluationMode,omitempty" yaml:"evaluationMode,omitempty"`
}
