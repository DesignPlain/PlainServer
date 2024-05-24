package kms

type Alias struct {
	// Identifier for the key for which the alias is for, can be either an ARN or key_id.
	TargetKeyId string `json:"targetKeyId,omitempty" yaml:"targetKeyId,omitempty"`

	// The display name of the alias. The name must start with the word "alias" followed by a forward slash (alias/)
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Creates an unique alias beginning with the specified prefix.
	   The name must start with the word "alias" followed by a forward slash (alias/).  Conflicts with `name`.
	*/
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`
}
