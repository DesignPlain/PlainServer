package types

type Lakeformation_PermissionsLfTagPolicyExpression struct {
	// The key-name of an LF-Tag.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// A list of possible values of an LF-Tag.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
