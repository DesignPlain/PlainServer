package types

type Medialive_ChannelEncoderSettingsOutputGroupOutputOutputSettingsUdpOutputSettingsFecOutputSettings struct {
	// The height of the FEC protection matrix.
	ColumnDepth int `json:"columnDepth,omitempty" yaml:"columnDepth,omitempty"`

	// Enables column only or column and row based FEC.
	IncludeFec string `json:"includeFec,omitempty" yaml:"includeFec,omitempty"`

	// The width of the FEC protection matrix.
	RowLength int `json:"rowLength,omitempty" yaml:"rowLength,omitempty"`
}
