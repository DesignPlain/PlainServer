package types

type Container_AttachedClusterLoggingConfigComponentConfig struct {
	/*
	   The components to be enabled.
	   Each value may be one of: `SYSTEM_COMPONENTS`, `WORKLOADS`.
	*/
	EnableComponents []string `json:"enableComponents,omitempty" yaml:"enableComponents,omitempty"`
}
