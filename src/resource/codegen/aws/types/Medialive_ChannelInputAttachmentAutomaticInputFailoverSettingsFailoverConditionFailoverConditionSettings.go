package types

type Medialive_ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettings struct {
	// MediaLive will perform a failover if content is considered black for the specified period. See Video Black Failover Settings for more details.
	VideoBlackSettings Medialive_ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsVideoBlackSettings `json:"videoBlackSettings,omitempty" yaml:"videoBlackSettings,omitempty"`

	// MediaLive will perform a failover if the specified audio selector is silent for the specified period. See Audio Silence Failover Settings for more details.
	AudioSilenceSettings Medialive_ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsAudioSilenceSettings `json:"audioSilenceSettings,omitempty" yaml:"audioSilenceSettings,omitempty"`

	// MediaLive will perform a failover if content is not detected in this input for the specified period. See Input Loss Failover Settings for more details.
	InputLossSettings Medialive_ChannelInputAttachmentAutomaticInputFailoverSettingsFailoverConditionFailoverConditionSettingsInputLossSettings `json:"inputLossSettings,omitempty" yaml:"inputLossSettings,omitempty"`
}
