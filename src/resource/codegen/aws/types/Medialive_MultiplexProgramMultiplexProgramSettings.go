package types

type Medialive_MultiplexProgramMultiplexProgramSettings struct {
	// Enum for preferred channel pipeline. Options are `CURRENTLY_ACTIVE`, `PIPELINE_0`, or `PIPELINE_1`.
	PreferredChannelPipeline string `json:"preferredChannelPipeline,omitempty" yaml:"preferredChannelPipeline,omitempty"`

	// Unique program number.
	ProgramNumber int `json:"programNumber,omitempty" yaml:"programNumber,omitempty"`

	// Service Descriptor. See Service Descriptor for more details.
	ServiceDescriptor Medialive_MultiplexProgramMultiplexProgramSettingsServiceDescriptor `json:"serviceDescriptor,omitempty" yaml:"serviceDescriptor,omitempty"`

	// Video settings. See Video Settings for more details.
	VideoSettings Medialive_MultiplexProgramMultiplexProgramSettingsVideoSettings `json:"videoSettings,omitempty" yaml:"videoSettings,omitempty"`
}
