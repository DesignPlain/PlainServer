package types

type Quicksight_getThemeConfiguration struct {
	// Color properties that apply to chart data colors. See data_color_palette.
	DataColorPalettes []Quicksight_getThemeConfigurationDataColorPalette `json:"dataColorPalettes,omitempty" yaml:"dataColorPalettes,omitempty"`

	// Display options related to sheets. See sheet.
	Sheets []Quicksight_getThemeConfigurationSheet `json:"sheets,omitempty" yaml:"sheets,omitempty"`

	// Determines the typography options. See typography.
	Typographies []Quicksight_getThemeConfigurationTypography `json:"typographies,omitempty" yaml:"typographies,omitempty"`

	// Color properties that apply to the UI and to charts, excluding the colors that apply to data. See ui_color_palette.
	UiColorPalettes []Quicksight_getThemeConfigurationUiColorPalette `json:"uiColorPalettes,omitempty" yaml:"uiColorPalettes,omitempty"`
}
