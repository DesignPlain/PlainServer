package securityhub

type ConfigurationPolicyAssociation struct {
	// The identifier of the target account, organizational unit, or the root to associate with the specified configuration.
	TargetId string `json:"targetId,omitempty" yaml:"targetId,omitempty"`

	// The universally unique identifier (UUID) of the configuration policy.
	PolicyId string `json:"policyId,omitempty" yaml:"policyId,omitempty"`
}
