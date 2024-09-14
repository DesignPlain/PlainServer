package types

type Diagflow_CxTestCaseTestCaseConversationTurnVirtualAgentOutput struct {
	/*
	   The [Page](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.flows.pages#Page) on which the utterance was spoken.
	   Structure is documented below.
	*/
	CurrentPage Diagflow_CxTestCaseTestCaseConversationTurnVirtualAgentOutputCurrentPage `json:"currentPage,omitempty" yaml:"currentPage,omitempty"`

	// The session parameters available to the bot at this point.
	SessionParameters string `json:"sessionParameters,omitempty" yaml:"sessionParameters,omitempty"`

	/*
	   The text responses from the agent for the turn.
	   Structure is documented below.
	*/
	TextResponses []Diagflow_CxTestCaseTestCaseConversationTurnVirtualAgentOutputTextResponse `json:"textResponses,omitempty" yaml:"textResponses,omitempty"`

	/*
	   The [Intent](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.intents#Intent) that triggered the response.
	   Structure is documented below.
	*/
	TriggeredIntent Diagflow_CxTestCaseTestCaseConversationTurnVirtualAgentOutputTriggeredIntent `json:"triggeredIntent,omitempty" yaml:"triggeredIntent,omitempty"`
}
