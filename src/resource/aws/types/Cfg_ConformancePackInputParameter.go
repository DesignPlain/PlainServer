package types

type Cfg_ConformancePackInputParameter struct {
	// The input key.
	ParameterName string `json:"parameterName,omitempty" yaml:"parameterName,omitempty"`

	// The input value.
	ParameterValue string `json:"parameterValue,omitempty" yaml:"parameterValue,omitempty"`
}
