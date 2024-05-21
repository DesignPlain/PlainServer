package types

type Container_NodePoolNodeConfigSandboxConfig struct {
	// Type of the sandbox to use for the node (e.g. 'gvisor')
	SandboxType string `json:"sandboxType,omitempty" yaml:"sandboxType,omitempty"`
}
