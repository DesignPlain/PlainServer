package types

type Medialive_ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsAudioSilenceSettings struct {
	// The name of the audio selector used as the source for this AudioDescription.
	AudioSelectorName string `json:"audioSelectorName,omitempty" yaml:"audioSelectorName,omitempty"`

	// The amount of time (in milliseconds) that the active input must be silent before automatic input failover occurs. Silence is defined as audio loss or audio quieter than -50 dBFS.
	AudioSilenceThresholdMsec int `json:"audioSilenceThresholdMsec,omitempty" yaml:"audioSilenceThresholdMsec,omitempty"`
}
