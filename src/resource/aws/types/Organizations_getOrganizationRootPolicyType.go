package types

type Organizations_getOrganizationRootPolicyType struct {
	// The status of the policy type as it relates to the associated root
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	//
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
