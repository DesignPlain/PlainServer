package types

type Quicksight_ThemeConfigurationSheet struct {
	// The display options for tiles. See tile.
	Tile Quicksight_ThemeConfigurationSheetTile `json:"tile,omitempty" yaml:"tile,omitempty"`

	// The layout options for tiles. See tile_layout.
	TileLayout Quicksight_ThemeConfigurationSheetTileLayout `json:"tileLayout,omitempty" yaml:"tileLayout,omitempty"`
}
