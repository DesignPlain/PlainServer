package types

type Firebase_ExtensionsInstanceRuntimeDataProcessingState struct {
	/*
	   Details about the processing. e.g. This could include the type of
	   processing in progress or it could list errors or failures.
	   This information will be shown in the console on the detail page
	   for the extension instance.
	*/
	DetailMessage string `json:"detailMessage,omitempty" yaml:"detailMessage,omitempty"`

	// The processing state of the extension instance.
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
