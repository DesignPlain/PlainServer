package types

type Bedrock_GuardrailTopicPolicyConfigTopicsConfig struct {
	// Definition of topic in topic policy.
	Definition string `json:"definition,omitempty" yaml:"definition,omitempty"`

	// List of text examples.
	Examples []string `json:"examples,omitempty" yaml:"examples,omitempty"`

	// Name of topic in topic policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Type of topic in a policy.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
