package types

type Iam_RoleInlinePolicy struct {
	// Name of the role policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Policy document as a JSON formatted string.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
