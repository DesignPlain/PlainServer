package types

type Firebase_ExtensionsInstanceRuntimeDataFatalError struct {
	/*
	   The error message. This is set by the extension developer to give
	   more detail on why the extension is unusable and must be re-installed
	   or reconfigured.
	*/
	ErrorMessage string `json:"errorMessage,omitempty" yaml:"errorMessage,omitempty"`
}
