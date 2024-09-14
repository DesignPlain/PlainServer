package types

type Quicksight_getThemeConfigurationTypography struct {
	// Determines the list of font families. Maximum number of 5 items. See font_families.
	FontFamilies []Quicksight_getThemeConfigurationTypographyFontFamily `json:"fontFamilies,omitempty" yaml:"fontFamilies,omitempty"`
}
