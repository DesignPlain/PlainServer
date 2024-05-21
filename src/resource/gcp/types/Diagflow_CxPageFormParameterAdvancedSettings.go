package types

type Diagflow_CxPageFormParameterAdvancedSettings struct {
	/*
	   Define behaviors for DTMF (dual tone multi frequency). DTMF settings does not override each other. DTMF settings set at different levels define DTMF detections running in parallel. Exposed at the following levels:
	   - Agent level
	   - Flow level
	   - Page level
	   - Parameter level
	   Structure is documented below.
	*/
	DtmfSettings Diagflow_CxPageFormParameterAdvancedSettingsDtmfSettings `json:"dtmfSettings,omitempty" yaml:"dtmfSettings,omitempty"`
}
