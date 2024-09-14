package types

type Medialive_ChannelInputAttachmentInputSettingsAudioSelectorSelectorSettingsAudioTrackSelection struct {
	// Configure decoding options for Dolby E streams - these should be Dolby E frames carried in PCM streams tagged with SMPTE-337. See Dolby E Decode for more details.
	DolbyEDecode Medialive_ChannelInputAttachmentInputSettingsAudioSelectorSelectorSettingsAudioTrackSelectionDolbyEDecode `json:"dolbyEDecode,omitempty" yaml:"dolbyEDecode,omitempty"`

	// Selects one or more unique audio tracks from within a source. See Audio Tracks for more details.
	Tracks []Medialive_ChannelInputAttachmentInputSettingsAudioSelectorSelectorSettingsAudioTrackSelectionTrack `json:"tracks,omitempty" yaml:"tracks,omitempty"`
}
