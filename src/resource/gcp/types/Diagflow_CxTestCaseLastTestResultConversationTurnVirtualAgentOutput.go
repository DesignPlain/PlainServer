package types

type Diagflow_CxTestCaseLastTestResultConversationTurnVirtualAgentOutput struct {
	/*
	   The [Page](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.flows.pages#Page) on which the utterance was spoken.
	   Structure is documented below.
	*/
	CurrentPage Diagflow_CxTestCaseLastTestResultConversationTurnVirtualAgentOutputCurrentPage `json:"currentPage,omitempty" yaml:"currentPage,omitempty"`

	/*
	   The list of differences between the original run and the replay for this output, if any.
	   Structure is documented below.
	*/
	Differences []Diagflow_CxTestCaseLastTestResultConversationTurnVirtualAgentOutputDifference `json:"differences,omitempty" yaml:"differences,omitempty"`

	// The session parameters available to the bot at this point.
	SessionParameters string `json:"sessionParameters,omitempty" yaml:"sessionParameters,omitempty"`

	/*
	   Response error from the agent in the test result. If set, other output is empty.
	   Structure is documented below.
	*/
	Status Diagflow_CxTestCaseLastTestResultConversationTurnVirtualAgentOutputStatus `json:"status,omitempty" yaml:"status,omitempty"`

	/*
	   The text responses from the agent for the turn.
	   Structure is documented below.
	*/
	TextResponses []Diagflow_CxTestCaseLastTestResultConversationTurnVirtualAgentOutputTextResponse `json:"textResponses,omitempty" yaml:"textResponses,omitempty"`

	/*
	   The [Intent](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.intents#Intent) that triggered the response.
	   Structure is documented below.
	*/
	TriggeredIntent Diagflow_CxTestCaseLastTestResultConversationTurnVirtualAgentOutputTriggeredIntent `json:"triggeredIntent,omitempty" yaml:"triggeredIntent,omitempty"`
}
