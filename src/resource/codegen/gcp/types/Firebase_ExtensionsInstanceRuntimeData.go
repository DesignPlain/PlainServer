package types

type Firebase_ExtensionsInstanceRuntimeData struct {
	/*
	   The fatal error state for the extension instance
	   Structure is documented below.
	*/
	FatalError Firebase_ExtensionsInstanceRuntimeDataFatalError `json:"fatalError,omitempty" yaml:"fatalError,omitempty"`

	/*
	   The processing state for the extension instance
	   Structure is documented below.
	*/
	ProcessingState Firebase_ExtensionsInstanceRuntimeDataProcessingState `json:"processingState,omitempty" yaml:"processingState,omitempty"`

	// The time of the last state update.
	StateUpdateTime string `json:"stateUpdateTime,omitempty" yaml:"stateUpdateTime,omitempty"`
}
