package types

type Diagflow_CxTestCaseLastTestResultConversationTurnVirtualAgentOutputDifference struct {
	// A human readable description of the diff, showing the actual output vs expected output.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The type of diff.
	   - INTENT: The intent.
	   - PAGE: The page.
	   - PARAMETERS: The parameters.
	   - UTTERANCE: The message utterance.
	   - FLOW: The flow.
	   Possible values are: `INTENT`, `PAGE`, `PARAMETERS`, `UTTERANCE`, `FLOW`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
