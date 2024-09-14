package types

type Bedrock_AgentAgentAliasRoutingConfiguration struct {
	// Version of the agent with which the alias is associated.
	AgentVersion string `json:"agentVersion,omitempty" yaml:"agentVersion,omitempty"`

	// ARN of the Provisioned Throughput assigned to the agent alias.
	ProvisionedThroughput string `json:"provisionedThroughput,omitempty" yaml:"provisionedThroughput,omitempty"`
}
