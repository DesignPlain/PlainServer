package mediastore

type ContainerPolicy struct {
	// The name of the container.
	ContainerName string `json:"containerName,omitempty" yaml:"containerName,omitempty"`

	// The contents of the policy.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
