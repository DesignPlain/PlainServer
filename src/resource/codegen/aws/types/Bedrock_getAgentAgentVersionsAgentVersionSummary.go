package types

type Bedrock_getAgentAgentVersionsAgentVersionSummary struct {
	//
	GuardrailConfigurations []Bedrock_getAgentAgentVersionsAgentVersionSummaryGuardrailConfiguration `json:"guardrailConfigurations,omitempty" yaml:"guardrailConfigurations,omitempty"`

	// Time at which the version was last updated.
	UpdatedAt string `json:"updatedAt,omitempty" yaml:"updatedAt,omitempty"`

	// Name of agent to which the version belongs.
	AgentName string `json:"agentName,omitempty" yaml:"agentName,omitempty"`

	// Status of the agent to which the version belongs.
	AgentStatus string `json:"agentStatus,omitempty" yaml:"agentStatus,omitempty"`

	// Version of the agent.
	AgentVersion string `json:"agentVersion,omitempty" yaml:"agentVersion,omitempty"`

	// Time at which the version was created.
	CreatedAt string `json:"createdAt,omitempty" yaml:"createdAt,omitempty"`

	/*
	   Description of the version of the agent.
	   - `GuardrailConfiguration` - Details aout the guardrail associated with the agent. See Guardrail Configuration
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
