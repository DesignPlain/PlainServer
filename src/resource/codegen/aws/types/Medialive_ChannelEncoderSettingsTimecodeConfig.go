package types

type Medialive_ChannelEncoderSettingsTimecodeConfig struct {
	// The source for the timecode that will be associated with the events outputs.
	Source string `json:"source,omitempty" yaml:"source,omitempty"`

	// Threshold in frames beyond which output timecode is resynchronized to the input timecode.
	SyncThreshold int `json:"syncThreshold,omitempty" yaml:"syncThreshold,omitempty"`
}
