package types

type Quicksight_ThemeConfigurationUiColorPalette struct {
	// Color (hexadecimal) that applies to the names of fields that are identified as dimensions.
	Dimension string `json:"dimension,omitempty" yaml:"dimension,omitempty"`

	// Color (hexadecimal) that applies to any text or other elements that appear over the measure color.
	MeasureForeground string `json:"measureForeground,omitempty" yaml:"measureForeground,omitempty"`

	// Color (hexadecimal) that applies to success messages, for example the check mark for a successful download.
	Success string `json:"success,omitempty" yaml:"success,omitempty"`

	// Color (hexadecimal) that applies to warning and informational messages.
	Warning string `json:"warning,omitempty" yaml:"warning,omitempty"`

	// Color (hexadecimal) that applies to any text or other elements that appear over the warning color.
	WarningForeground string `json:"warningForeground,omitempty" yaml:"warningForeground,omitempty"`

	// Color (hexadecimal) that applies to error messages.
	Danger string `json:"danger,omitempty" yaml:"danger,omitempty"`

	// Color (hexadecimal) that applies to any sheet title, sheet control text, or UI that appears over the secondary background.
	SecondaryForeground string `json:"secondaryForeground,omitempty" yaml:"secondaryForeground,omitempty"`

	// Color (hexadecimal) that applies to selected states and buttons.
	Accent string `json:"accent,omitempty" yaml:"accent,omitempty"`

	// Color (hexadecimal) that applies to any text or other elements that appear over the dimension color.
	DimensionForeground string `json:"dimensionForeground,omitempty" yaml:"dimensionForeground,omitempty"`

	// Color (hexadecimal) of text and other foreground elements that appear over the primary background regions, such as grid lines, borders, table banding, icons, and so on.
	PrimaryForeground string `json:"primaryForeground,omitempty" yaml:"primaryForeground,omitempty"`

	// Color (hexadecimal) that applies to any text or other elements that appear over the success color.
	SuccessForeground string `json:"successForeground,omitempty" yaml:"successForeground,omitempty"`

	// Color (hexadecimal) that applies to any text or other elements that appear over the accent color.
	AccentForeground string `json:"accentForeground,omitempty" yaml:"accentForeground,omitempty"`

	// Color (hexadecimal) that applies to any text or other elements that appear over the error color.
	DangerForeground string `json:"dangerForeground,omitempty" yaml:"dangerForeground,omitempty"`

	// Color (hexadecimal) that applies to the names of fields that are identified as measures.
	Measure string `json:"measure,omitempty" yaml:"measure,omitempty"`

	// Color (hexadecimal) that applies to visuals and other high emphasis UI.
	PrimaryBackground string `json:"primaryBackground,omitempty" yaml:"primaryBackground,omitempty"`

	// Color (hexadecimal) that applies to the sheet background and sheet controls.
	SecondaryBackground string `json:"secondaryBackground,omitempty" yaml:"secondaryBackground,omitempty"`
}
