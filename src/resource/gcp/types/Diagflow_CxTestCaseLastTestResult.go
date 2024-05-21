package types

type Diagflow_CxTestCaseLastTestResult struct {
	/*
	   The conversation turns uttered during the test case replay in chronological order.
	   Structure is documented below.
	*/
	ConversationTurns []Diagflow_CxTestCaseLastTestResultConversationTurn `json:"conversationTurns,omitempty" yaml:"conversationTurns,omitempty"`

	// Environment where the test was run. If not set, it indicates the draft environment.
	Environment string `json:"environment,omitempty" yaml:"environment,omitempty"`

	/*
	   The unique identifier of the page.
	   Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Whether the test case passed in the agent environment.
	   - PASSED: The test passed.
	   - FAILED: The test did not pass.
	   Possible values are: `PASSED`, `FAILED`.
	*/
	TestResult string `json:"testResult,omitempty" yaml:"testResult,omitempty"`

	// The time that the test was run. A timestamp in RFC3339 text format.
	TestTime string `json:"testTime,omitempty" yaml:"testTime,omitempty"`
}
