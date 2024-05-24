package types

type Ses_EventDestinationCloudwatchDestination struct {
	// The default value for the event
	DefaultValue string `json:"defaultValue,omitempty" yaml:"defaultValue,omitempty"`

	// The name for the dimension
	DimensionName string `json:"dimensionName,omitempty" yaml:"dimensionName,omitempty"`

	// The source for the value. May be any of `"messageTag"`, `"emailHeader"` or `"linkTag"`.
	ValueSource string `json:"valueSource,omitempty" yaml:"valueSource,omitempty"`
}
