package types

type Diagflow_CxTestCaseTestCaseConversationTurnVirtualAgentOutputCurrentPage struct {
	/*
	   (Output)
	   The human-readable name of the page, unique within the flow.
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The unique identifier of the page.
	   Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/pages/<Page ID>.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
