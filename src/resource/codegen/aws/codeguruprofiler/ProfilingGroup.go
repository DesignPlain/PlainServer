package codeguruprofiler

import types "libds/aws/types"

type ProfilingGroup struct {
	// Specifies whether profiling is enabled or disabled for the created profiling. See Agent Orchestration Config for more details.
	AgentOrchestrationConfig types.Codeguruprofiler_ProfilingGroupAgentOrchestrationConfig `json:"agentOrchestrationConfig,omitempty" yaml:"agentOrchestrationConfig,omitempty"`

	// Compute platform of the profiling group.
	ComputePlatform string `json:"computePlatform,omitempty" yaml:"computePlatform,omitempty"`

	/*
	   Name of the profiling group.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Map of tags assigned to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
