package schemas

type RegistryPolicy struct {
	// Resource Policy for EventBridge Schema Registry
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`

	// Name of EventBridge Schema Registry
	RegistryName string `json:"registryName,omitempty" yaml:"registryName,omitempty"`
}
