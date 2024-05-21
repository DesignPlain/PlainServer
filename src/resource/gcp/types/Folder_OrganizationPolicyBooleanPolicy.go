package types

type Folder_OrganizationPolicyBooleanPolicy struct {
	// If true, then the Policy is enforced. If false, then any configuration is acceptable.
	Enforced bool `json:"enforced,omitempty" yaml:"enforced,omitempty"`
}
