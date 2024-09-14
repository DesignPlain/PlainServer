package types

type Projects_getOrganizationPolicyListPolicyDeny struct {
	// The policy can define specific values that are allowed or denied.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	// The policy allows or denies all values.
	All bool `json:"all,omitempty" yaml:"all,omitempty"`
}
