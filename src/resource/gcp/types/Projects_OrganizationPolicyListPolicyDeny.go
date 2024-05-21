package types

type Projects_OrganizationPolicyListPolicyDeny struct {
	// The policy allows or denies all values.
	All bool `json:"all,omitempty" yaml:"all,omitempty"`

	// The policy can define specific values that are allowed or denied.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
