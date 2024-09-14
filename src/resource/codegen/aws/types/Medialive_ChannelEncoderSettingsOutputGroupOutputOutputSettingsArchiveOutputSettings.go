package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettings struct {
	// Settings specific to the container type of the file. See Container Settings for more details.
	ContainerSettings Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsArchiveOutputSettingsContainerSettings `json:"containerSettings,omitempty" yaml:"containerSettings,omitempty"`

	// Output file extension.
	Extension string `json:"extension,omitempty" yaml:"extension,omitempty"`

	// String concatenated to the end of the destination filename. Required for multiple outputs of the same type.
	NameModifier string `json:"nameModifier,omitempty" yaml:"nameModifier,omitempty"`
}
