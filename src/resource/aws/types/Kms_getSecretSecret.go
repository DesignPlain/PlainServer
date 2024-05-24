package types

type Kms_getSecretSecret struct {
	//
	GrantTokens []string `json:"grantTokens,omitempty" yaml:"grantTokens,omitempty"`

	//
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Payload string `json:"payload,omitempty" yaml:"payload,omitempty"`

	//
	Context map[string]string `json:"context,omitempty" yaml:"context,omitempty"`
}
