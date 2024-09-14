package types

type Dlm_LifecyclePolicyPolicyDetailsEventSource struct {
	//
	Parameters Dlm_LifecyclePolicyPolicyDetailsEventSourceParameters `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// The source of the event. Currently only managed CloudWatch Events rules are supported. Valid values are `MANAGED_CWE`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
