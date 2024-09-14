package bedrock

import types "libds/aws/types"

type AgentAgentAlias struct {
	// Name of the alias.
	AgentAliasName string `json:"agentAliasName,omitempty" yaml:"agentAliasName,omitempty"`

	/*
	   Identifier of the agent to create an alias for.

	   The following arguments are optional:
	*/
	AgentId string `json:"agentId,omitempty" yaml:"agentId,omitempty"`

	// Description of the alias.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Details about the routing configuration of the alias. See `routing_configuration` Block for details.
	RoutingConfigurations []types.Bedrock_AgentAgentAliasRoutingConfiguration `json:"routingConfigurations,omitempty" yaml:"routingConfigurations,omitempty"`

	// Map of tags assigned to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Bedrock_AgentAgentAliasTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
