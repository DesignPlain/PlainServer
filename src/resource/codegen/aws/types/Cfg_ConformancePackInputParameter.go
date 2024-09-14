package types

type Cfg_ConformancePackInputParameter struct {
	// The input value.
	ParameterValue string `json:"parameterValue,omitempty" yaml:"parameterValue,omitempty"`

	// The input key.
	ParameterName string `json:"parameterName,omitempty" yaml:"parameterName,omitempty"`
}
