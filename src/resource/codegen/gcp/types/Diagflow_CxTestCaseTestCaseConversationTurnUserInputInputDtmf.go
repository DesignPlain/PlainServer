package types

type Diagflow_CxTestCaseTestCaseConversationTurnUserInputInputDtmf struct {
	// The dtmf digits.
	Digits string `json:"digits,omitempty" yaml:"digits,omitempty"`

	// The finish digit (if any).
	FinishDigit string `json:"finishDigit,omitempty" yaml:"finishDigit,omitempty"`
}
