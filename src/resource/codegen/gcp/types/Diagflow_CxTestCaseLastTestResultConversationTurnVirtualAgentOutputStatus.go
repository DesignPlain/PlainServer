package types

type Diagflow_CxTestCaseLastTestResultConversationTurnVirtualAgentOutputStatus struct {
	// A JSON encoded list of messages that carry the error details.
	Details string `json:"details,omitempty" yaml:"details,omitempty"`

	// A developer-facing error message.
	Message string `json:"message,omitempty" yaml:"message,omitempty"`

	// The status code, which should be an enum value of google.rpc.Code.
	Code int `json:"code,omitempty" yaml:"code,omitempty"`
}
