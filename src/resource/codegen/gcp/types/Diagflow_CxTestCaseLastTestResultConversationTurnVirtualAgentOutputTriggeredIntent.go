package types

type Diagflow_CxTestCaseLastTestResultConversationTurnVirtualAgentOutputTriggeredIntent struct {
	/*
	   (Output)
	   The human-readable name of the intent, unique within the agent.
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The unique identifier of the intent.
	   Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/intents/<Intent ID>.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
