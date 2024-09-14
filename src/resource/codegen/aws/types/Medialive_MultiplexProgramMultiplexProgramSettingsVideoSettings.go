package types

type Medialive_MultiplexProgramMultiplexProgramSettingsVideoSettings struct {
	// Constant bitrate value.
	ConstantBitrate int `json:"constantBitrate,omitempty" yaml:"constantBitrate,omitempty"`

	// Statmux settings. See Statmux Settings for more details.
	StatmuxSettings Medialive_MultiplexProgramMultiplexProgramSettingsVideoSettingsStatmuxSettings `json:"statmuxSettings,omitempty" yaml:"statmuxSettings,omitempty"`
}
