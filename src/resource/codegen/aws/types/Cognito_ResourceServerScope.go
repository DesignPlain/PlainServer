package types

type Cognito_ResourceServerScope struct {
	// The scope description.
	ScopeDescription string `json:"scopeDescription,omitempty" yaml:"scopeDescription,omitempty"`

	// The scope name.
	ScopeName string `json:"scopeName,omitempty" yaml:"scopeName,omitempty"`
}
