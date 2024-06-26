package types

type Diagflow_CxAgentAdvancedSettings struct {
	/*
	   If present, incoming audio is exported by Dialogflow to the configured Google Cloud Storage destination. Exposed at the following levels:
	   - Agent level
	   - Flow level
	   Structure is documented below.
	*/
	AudioExportGcsDestination Diagflow_CxAgentAdvancedSettingsAudioExportGcsDestination `json:"audioExportGcsDestination,omitempty" yaml:"audioExportGcsDestination,omitempty"`

	/*
	   Define behaviors for DTMF (dual tone multi frequency). DTMF settings does not override each other. DTMF settings set at different levels define DTMF detections running in parallel. Exposed at the following levels:
	   - Agent level
	   - Flow level
	   - Page level
	   - Parameter level
	   Structure is documented below.
	*/
	DtmfSettings Diagflow_CxAgentAdvancedSettingsDtmfSettings `json:"dtmfSettings,omitempty" yaml:"dtmfSettings,omitempty"`
}
