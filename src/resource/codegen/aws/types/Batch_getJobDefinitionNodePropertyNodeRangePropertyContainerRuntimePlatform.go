package types

type Batch_getJobDefinitionNodePropertyNodeRangePropertyContainerRuntimePlatform struct {
	// The vCPU architecture. The default value is X86_64. Valid values are X86_64 and ARM64.
	CpuArchitecture string `json:"cpuArchitecture,omitempty" yaml:"cpuArchitecture,omitempty"`

	// The operating system for the compute environment. V
	OperatingSystemFamily string `json:"operatingSystemFamily,omitempty" yaml:"operatingSystemFamily,omitempty"`
}
