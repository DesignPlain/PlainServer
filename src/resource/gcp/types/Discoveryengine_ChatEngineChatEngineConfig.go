package types

type Discoveryengine_ChatEngineChatEngineConfig struct {
	/*
	   The configuration to generate the Dialogflow agent that is associated to this Engine.
	   Structure is documented below.
	*/
	AgentCreationConfig Discoveryengine_ChatEngineChatEngineConfigAgentCreationConfig `json:"agentCreationConfig,omitempty" yaml:"agentCreationConfig,omitempty"`
}
