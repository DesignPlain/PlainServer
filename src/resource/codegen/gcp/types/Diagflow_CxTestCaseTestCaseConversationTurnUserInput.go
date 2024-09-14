package types

type Diagflow_CxTestCaseTestCaseConversationTurnUserInput struct {
	// Parameters that need to be injected into the conversation during intent detection.
	InjectedParameters string `json:"injectedParameters,omitempty" yaml:"injectedParameters,omitempty"`

	/*
	   User input. Supports text input, event input, dtmf input in the test case.
	   Structure is documented below.
	*/
	Input Diagflow_CxTestCaseTestCaseConversationTurnUserInputInput `json:"input,omitempty" yaml:"input,omitempty"`

	// If webhooks should be allowed to trigger in response to the user utterance. Often if parameters are injected, webhooks should not be enabled.
	IsWebhookEnabled bool `json:"isWebhookEnabled,omitempty" yaml:"isWebhookEnabled,omitempty"`

	// Whether sentiment analysis is enabled.
	EnableSentimentAnalysis bool `json:"enableSentimentAnalysis,omitempty" yaml:"enableSentimentAnalysis,omitempty"`
}
