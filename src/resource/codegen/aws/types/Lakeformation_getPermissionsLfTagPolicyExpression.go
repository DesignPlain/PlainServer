package types

type Lakeformation_getPermissionsLfTagPolicyExpression struct {
	// Key-name of an LF-Tag.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// List of possible values of an LF-Tag.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
