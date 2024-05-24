package types

type Quicksight_getThemeConfigurationSheet struct {
	// The layout options for tiles. See tile_layout.
	TileLayouts []Quicksight_getThemeConfigurationSheetTileLayout `json:"tileLayouts,omitempty" yaml:"tileLayouts,omitempty"`

	// The display options for tiles. See tile.
	Tiles []Quicksight_getThemeConfigurationSheetTile `json:"tiles,omitempty" yaml:"tiles,omitempty"`
}
