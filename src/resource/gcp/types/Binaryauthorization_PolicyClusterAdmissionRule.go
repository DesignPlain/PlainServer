package types

type Binaryauthorization_PolicyClusterAdmissionRule struct {
	/*
	   The action when a pod creation is denied by the admission rule.
	   Possible values are: `ENFORCED_BLOCK_AND_AUDIT_LOG`, `DRYRUN_AUDIT_LOG_ONLY`.
	*/
	EnforcementMode string `json:"enforcementMode,omitempty" yaml:"enforcementMode,omitempty"`

	/*
	   How this admission rule will be evaluated.
	   Possible values are: `ALWAYS_ALLOW`, `REQUIRE_ATTESTATION`, `ALWAYS_DENY`.
	*/
	EvaluationMode string `json:"evaluationMode,omitempty" yaml:"evaluationMode,omitempty"`

	/*
	   The resource names of the attestors that must attest to a
	   container image. If the attestor is in a different project from the
	   policy, it should be specified in the format `projects/-/attestors/-`.
	   Each attestor must exist before a policy can reference it. To add an
	   attestor to a policy the principal issuing the policy change
	   request must be able to read the attestor resource.
	   Note: this field must be non-empty when the evaluation_mode field
	   specifies REQUIRE_ATTESTATION, otherwise it must be empty.
	*/
	RequireAttestationsBies []string `json:"requireAttestationsBies,omitempty" yaml:"requireAttestationsBies,omitempty"`

	// The identifier for this object. Format specified above.
	Cluster string `json:"cluster,omitempty" yaml:"cluster,omitempty"`
}
