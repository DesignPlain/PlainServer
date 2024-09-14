package ecr

type RegistryPolicy struct {
	// The policy document. This is a JSON formatted string.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
