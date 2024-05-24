package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSetting struct {
	// Number of seconds to write to archive file before closing and starting a new one.
	RolloverInterval int `json:"rolloverInterval,omitempty" yaml:"rolloverInterval,omitempty"`

	// Parameters that control the interactions with the CDN. See Archive CDN Settings for more details.
	ArchiveCdnSettings Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingArchiveCdnSettings `json:"archiveCdnSettings,omitempty" yaml:"archiveCdnSettings,omitempty"`

	// A director and base filename where archive files should be written. See Destination for more details.
	Destination Medialive_ChannelEncoderSettingsOutputGroupOutputGroupSettingsArchiveGroupSettingDestination `json:"destination,omitempty" yaml:"destination,omitempty"`
}
