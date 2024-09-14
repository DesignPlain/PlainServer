package types

type Quicksight_ThemeConfigurationDataColorPalette struct {
	// List of hexadecimal codes for the colors. Minimum of 8 items and maximum of 20 items.
	Colors []string `json:"colors,omitempty" yaml:"colors,omitempty"`

	// The hexadecimal code of a color that applies to charts where a lack of data is highlighted.
	EmptyFillColor string `json:"emptyFillColor,omitempty" yaml:"emptyFillColor,omitempty"`

	// The minimum and maximum hexadecimal codes that describe a color gradient. List of exactly 2 items.
	MinMaxGradients []string `json:"minMaxGradients,omitempty" yaml:"minMaxGradients,omitempty"`
}
