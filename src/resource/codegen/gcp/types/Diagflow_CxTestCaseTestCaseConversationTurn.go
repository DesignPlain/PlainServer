package types

type Diagflow_CxTestCaseTestCaseConversationTurn struct {
	/*
	   The user input.
	   Structure is documented below.
	*/
	UserInput Diagflow_CxTestCaseTestCaseConversationTurnUserInput `json:"userInput,omitempty" yaml:"userInput,omitempty"`

	/*
	   The virtual agent output.
	   Structure is documented below.
	*/
	VirtualAgentOutput Diagflow_CxTestCaseTestCaseConversationTurnVirtualAgentOutput `json:"virtualAgentOutput,omitempty" yaml:"virtualAgentOutput,omitempty"`
}
