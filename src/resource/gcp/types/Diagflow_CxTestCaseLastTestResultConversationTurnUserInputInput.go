package types

type Diagflow_CxTestCaseLastTestResultConversationTurnUserInputInput struct {
	/*
	   The language of the input. See [Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language) for a list of the currently supported language codes.
	   Note that queries in the same session do not necessarily need to specify the same language.
	*/
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	/*
	   The natural language text to be processed.
	   Structure is documented below.
	*/
	Text Diagflow_CxTestCaseLastTestResultConversationTurnUserInputInputText `json:"text,omitempty" yaml:"text,omitempty"`

	/*
	   The DTMF event to be handled.
	   Structure is documented below.
	*/
	Dtmf Diagflow_CxTestCaseLastTestResultConversationTurnUserInputInputDtmf `json:"dtmf,omitempty" yaml:"dtmf,omitempty"`

	/*
	   The event to be triggered.
	   Structure is documented below.
	*/
	Event Diagflow_CxTestCaseLastTestResultConversationTurnUserInputInputEvent `json:"event,omitempty" yaml:"event,omitempty"`
}
