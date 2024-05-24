package types

type Quicksight_ThemeConfiguration struct {
	// Color properties that apply to chart data colors. See data_color_palette.
	DataColorPalette Quicksight_ThemeConfigurationDataColorPalette `json:"dataColorPalette,omitempty" yaml:"dataColorPalette,omitempty"`

	// Display options related to sheets. See sheet.
	Sheet Quicksight_ThemeConfigurationSheet `json:"sheet,omitempty" yaml:"sheet,omitempty"`

	// Determines the typography options. See typography.
	Typography Quicksight_ThemeConfigurationTypography `json:"typography,omitempty" yaml:"typography,omitempty"`

	// Color properties that apply to the UI and to charts, excluding the colors that apply to data. See ui_color_palette.
	UiColorPalette Quicksight_ThemeConfigurationUiColorPalette `json:"uiColorPalette,omitempty" yaml:"uiColorPalette,omitempty"`
}
