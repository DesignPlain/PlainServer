package mediastore

type ContainerPolicy struct {
	// The contents of the policy.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// The name of the container.
	ContainerName string `json:"containerName,omitempty" yaml:"containerName,omitempty"`
}
