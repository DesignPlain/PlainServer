package types

type Diagflow_CxTestCaseLastTestResultConversationTurn struct {
	/*
	   The user input.
	   Structure is documented below.
	*/
	UserInput Diagflow_CxTestCaseLastTestResultConversationTurnUserInput `json:"userInput,omitempty" yaml:"userInput,omitempty"`

	/*
	   The virtual agent output.
	   Structure is documented below.
	*/
	VirtualAgentOutput Diagflow_CxTestCaseLastTestResultConversationTurnVirtualAgentOutput `json:"virtualAgentOutput,omitempty" yaml:"virtualAgentOutput,omitempty"`
}
