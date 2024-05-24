package types

type Lambda_FunctionEnvironment struct {
	// Map of environment variables that are accessible from the function code during execution. If provided at least one key must be present.
	Variables map[string]string `json:"variables,omitempty" yaml:"variables,omitempty"`
}
