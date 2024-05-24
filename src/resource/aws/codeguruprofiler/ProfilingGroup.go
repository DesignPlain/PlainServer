package codeguruprofiler

import types "DesignSphere_Server/src/resource/aws/types"

type ProfilingGroup struct {
	/*
	   Name of the profiling group.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags assigned to the WorkSpaces Connection Alias. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies whether profiling is enabled or disabled for the created profiling. See Agent Orchestration Config for more details.
	AgentOrchestrationConfig types.Codeguruprofiler_ProfilingGroupAgentOrchestrationConfig `json:"agentOrchestrationConfig,omitempty" yaml:"agentOrchestrationConfig,omitempty"`

	// Compute platform of the profiling group.
	ComputePlatform string `json:"computePlatform,omitempty" yaml:"computePlatform,omitempty"`
}
