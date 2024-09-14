package types

type Codebuild_ProjectEnvironmentEnvironmentVariable struct {
	// Environment variable's name or key.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Type of environment variable. Valid values: `PARAMETER_STORE`, `PLAINTEXT`, `SECRETS_MANAGER`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Environment variable's value.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
