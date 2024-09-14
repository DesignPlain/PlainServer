package types

type Medialive_ChannelEncoderSettingsOutputGroup struct {
	// Custom output group name defined by the user.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Settings associated with the output group. See Output Group Settings for more details.
	OutputGroupSettings Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettings `json:"outputGroupSettings,omitempty" yaml:"outputGroupSettings,omitempty"`

	// List of outputs. See Outputs for more details.
	Outputs []Medialive_ChannelEncoderSettingsOutputGroupOutput `json:"outputs,omitempty" yaml:"outputs,omitempty"`
}
