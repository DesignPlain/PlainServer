package types

type Medialive_ChannelInputAttachmentInputSettingsAudioSelectorSelectorSettingsAudioHlsRenditionSelection struct {
	// Specifies the NAME in the #EXT-X-MEDIA tag of the target HLS audio rendition.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies the GROUP-ID in the #EXT-X-MEDIA tag of the target HLS audio rendition.
	GroupId string `json:"groupId,omitempty" yaml:"groupId,omitempty"`
}
