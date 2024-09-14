package types

type Quicksight_getThemeConfigurationSheetTileLayout struct {
	// The gutter settings that apply between tiles. See gutter.
	Gutters []Quicksight_getThemeConfigurationSheetTileLayoutGutter `json:"gutters,omitempty" yaml:"gutters,omitempty"`

	// The margin settings that apply around the outside edge of sheets. See margin.
	Margins []Quicksight_getThemeConfigurationSheetTileLayoutMargin `json:"margins,omitempty" yaml:"margins,omitempty"`
}
