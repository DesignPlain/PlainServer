package types

type Medialive_ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsTeletextSourceSettings struct {
	// Optionally defines a region where TTML style captions will be displayed. See Caption Rectangle for more details.
	OutputRectangle Medialive_ChannelInputAttachmentInputSettingsCaptionSelectorSelectorSettingsTeletextSourceSettingsOutputRectangle `json:"outputRectangle,omitempty" yaml:"outputRectangle,omitempty"`

	// Specifies the teletext page number within the data stream from which to extract captions. Range of 0x100 (256) to 0x8FF (2303). Unused for passthrough. Should be specified as a hexadecimal string with no “0x” prefix.
	PageNumber string `json:"pageNumber,omitempty" yaml:"pageNumber,omitempty"`
}
