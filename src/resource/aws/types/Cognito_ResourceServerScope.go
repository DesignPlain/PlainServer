package types

type Cognito_ResourceServerScope struct {
	// The scope name.
	ScopeName string `json:"scopeName,omitempty" yaml:"scopeName,omitempty"`

	// The scope description.
	ScopeDescription string `json:"scopeDescription,omitempty" yaml:"scopeDescription,omitempty"`
}
