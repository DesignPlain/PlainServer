package types

type Iam_getPrincipalPolicySimulationResultMatchedStatement struct {
	// The type of the policy identified in source_policy_id.
	SourcePolicyType string `json:"sourcePolicyType,omitempty" yaml:"sourcePolicyType,omitempty"`

	// Identifier of one of the policies used as input to the simulation.
	SourcePolicyId string `json:"sourcePolicyId,omitempty" yaml:"sourcePolicyId,omitempty"`
}
