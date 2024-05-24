package types

type Quicksight_ThemeConfigurationSheetTileLayout struct {
	// The margin settings that apply around the outside edge of sheets. See margin.
	Margin Quicksight_ThemeConfigurationSheetTileLayoutMargin `json:"margin,omitempty" yaml:"margin,omitempty"`

	// The gutter settings that apply between tiles. See gutter.
	Gutter Quicksight_ThemeConfigurationSheetTileLayoutGutter `json:"gutter,omitempty" yaml:"gutter,omitempty"`
}
