package types

type Container_getClusterNodePoolNodeConfigSandboxConfig struct {
	// Type of the sandbox to use for the node (e.g. 'gvisor')
	SandboxType string `json:"sandboxType,omitempty" yaml:"sandboxType,omitempty"`
}
