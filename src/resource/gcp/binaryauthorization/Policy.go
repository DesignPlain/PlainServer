package binaryauthorization

import types "DesignSphere_Server/src/resource/gcp/types"

type Policy struct {
	/*
	   Per-cluster admission rules. An admission rule specifies either that
	   all container images used in a pod creation request must be attested
	   to by one or more attestors, that all pod creations will be allowed,
	   or that all pod creations will be denied. There can be at most one
	   admission rule per cluster spec.

	   Identifier format: `{{location}}.{{clusterId}}`.
	   A location is either a compute zone (e.g. `us-central1-a`) or a region
	   (e.g. `us-central1`).
	   Structure is documented below.
	*/
	ClusterAdmissionRules []types.Binaryauthorization_PolicyClusterAdmissionRule `json:"clusterAdmissionRules,omitempty" yaml:"clusterAdmissionRules,omitempty"`

	/*
	   Default admission rule for a cluster without a per-cluster admission
	   rule.
	   Structure is documented below.
	*/
	DefaultAdmissionRule types.Binaryauthorization_PolicyDefaultAdmissionRule `json:"defaultAdmissionRule,omitempty" yaml:"defaultAdmissionRule,omitempty"`

	// A descriptive comment.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Controls the evaluation of a Google-maintained global admission policy
	   for common system-level images. Images not covered by the global
	   policy will be subject to the project admission policy.
	   Possible values are: `ENABLE`, `DISABLE`.
	*/
	GlobalPolicyEvaluationMode string `json:"globalPolicyEvaluationMode,omitempty" yaml:"globalPolicyEvaluationMode,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A whitelist of image patterns to exclude from admission rules. If an
	   image's name matches a whitelist pattern, the image's admission
	   requests will always be permitted regardless of your admission rules.
	   Structure is documented below.
	*/
	AdmissionWhitelistPatterns []types.Binaryauthorization_PolicyAdmissionWhitelistPattern `json:"admissionWhitelistPatterns,omitempty" yaml:"admissionWhitelistPatterns,omitempty"`
}
