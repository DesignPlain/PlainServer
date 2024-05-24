package types

type Eks_AccessPolicyAssociationAccessScope struct {
	// The namespaces to which the access scope applies when type is namespace.
	Namespaces []string `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`

	// Valid values are `namespace` or `cluster`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
