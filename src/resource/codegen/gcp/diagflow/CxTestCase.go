package diagflow

import types "libds/gcp/types"

type CxTestCase struct {
	/*
	   The conversation turns uttered when the test case was created, in chronological order. These include the canonical set of agent utterances that should occur when the agent is working properly.
	   Structure is documented below.
	*/
	TestCaseConversationTurns []types.Diagflow_CxTestCaseTestCaseConversationTurn `json:"testCaseConversationTurns,omitempty" yaml:"testCaseConversationTurns,omitempty"`

	/*
	   Config for the test case.
	   Structure is documented below.
	*/
	TestConfig types.Diagflow_CxTestCaseTestConfig `json:"testConfig,omitempty" yaml:"testConfig,omitempty"`

	/*
	   The human-readable name of the test case, unique within the agent. Limit of 200 characters.


	   - - -
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Additional freeform notes about the test case. Limit of 400 characters.
	Notes string `json:"notes,omitempty" yaml:"notes,omitempty"`

	/*
	   The agent to create the test case for.
	   Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	/*
	   Tags are short descriptions that users may apply to test cases for organizational and filtering purposes.
	   Each tag should start with "#" and has a limit of 30 characters
	*/
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
