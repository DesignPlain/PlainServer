package types

type Diagflow_CxFlowAdvancedSettingsAudioExportGcsDestination struct {
	/*
	   The Google Cloud Storage URI for the exported objects. Whether a full object name, or just a prefix, its usage depends on the Dialogflow operation.
	   Format: gs://bucket/object-name-or-prefix
	*/
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`
}
