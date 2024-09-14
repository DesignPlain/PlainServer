package diagflow

import types "libds/gcp/types"

type CxEnvironment struct {
	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The Agent to create an Environment for.
	   Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	/*
	   A list of configurations for flow versions. You should include version configs for all flows that are reachable from [Start Flow][Agent.start_flow] in the agent. Otherwise, an error will be returned.
	   Structure is documented below.
	*/
	VersionConfigs []types.Diagflow_CxEnvironmentVersionConfig `json:"versionConfigs,omitempty" yaml:"versionConfigs,omitempty"`
}
