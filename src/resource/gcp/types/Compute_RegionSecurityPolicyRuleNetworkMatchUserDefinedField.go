package types

type Compute_RegionSecurityPolicyRuleNetworkMatchUserDefinedField struct {
	// Matching values of the field. Each element can be a 32-bit unsigned decimal or hexadecimal (starting with "0x") number (e.g. "64") or range (e.g. "0x400-0x7ff").
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`

	// Name of the user-defined field, as given in the definition.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
