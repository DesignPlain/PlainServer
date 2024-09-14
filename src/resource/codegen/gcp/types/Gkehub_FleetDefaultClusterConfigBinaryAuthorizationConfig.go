package types

type Gkehub_FleetDefaultClusterConfigBinaryAuthorizationConfig struct {
	/*
	   Mode of operation for binauthz policy evaluation.
	   Possible values are: `DISABLED`, `POLICY_BINDINGS`.
	*/
	EvaluationMode string `json:"evaluationMode,omitempty" yaml:"evaluationMode,omitempty"`

	/*
	   Binauthz policies that apply to this cluster.
	   Structure is documented below.
	*/
	PolicyBindings []Gkehub_FleetDefaultClusterConfigBinaryAuthorizationConfigPolicyBinding `json:"policyBindings,omitempty" yaml:"policyBindings,omitempty"`
}
