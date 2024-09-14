package bedrock

import types "libds/aws/types"

type AgentAgentKnowledgeBaseAssociation struct {
	// Description of what the agent should use the knowledge base for.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Unique identifier of the knowledge base to associate with the agent.
	KnowledgeBaseId string `json:"knowledgeBaseId,omitempty" yaml:"knowledgeBaseId,omitempty"`

	/*
	   Whether to use the knowledge base when sending an [InvokeAgent](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent-runtime_InvokeAgent.html) request. Valid values: `ENABLED`, `DISABLED`.

	   The following arguments are optional:
	*/
	KnowledgeBaseState string `json:"knowledgeBaseState,omitempty" yaml:"knowledgeBaseState,omitempty"`

	//
	Timeouts types.Bedrock_AgentAgentKnowledgeBaseAssociationTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Unique identifier of the agent with which you want to associate the knowledge base.
	AgentId string `json:"agentId,omitempty" yaml:"agentId,omitempty"`

	// Version of the agent with which you want to associate the knowledge base. Valid values: `DRAFT`.
	AgentVersion string `json:"agentVersion,omitempty" yaml:"agentVersion,omitempty"`
}
