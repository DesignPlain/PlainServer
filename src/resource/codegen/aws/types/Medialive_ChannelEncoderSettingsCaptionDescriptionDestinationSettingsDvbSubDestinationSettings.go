package types

type Medialive_ChannelEncoderSettingsCaptionDescriptionDestinationSettingsDvbSubDestinationSettings struct {
	// Specifies the color of the rectangle behind the captions. All burn-in and DVB-Sub font settings must match.
	BackgroundColor string `json:"backgroundColor,omitempty" yaml:"backgroundColor,omitempty"`

	// Specifies the opacity of the burned-in captions. 255 is opaque; 0 is transparent. All burn-in and DVB-Sub font settings must match.
	FontOpacity int `json:"fontOpacity,omitempty" yaml:"fontOpacity,omitempty"`

	// Font resolution in DPI (dots per inch); default is 96 dpi. All burn-in and DVB-Sub font settings must match.
	FontResolution int `json:"fontResolution,omitempty" yaml:"fontResolution,omitempty"`

	// Specifies the horizontal offset of the shadow relative to the captions in pixels. A value of -2 would result in a shadow offset 2 pixels to the left. All burn-in and DVB-Sub font settings must match.
	ShadowXOffset int `json:"shadowXOffset,omitempty" yaml:"shadowXOffset,omitempty"`

	// Specifies the vertical offset of the shadow relative to the captions in pixels. A value of -2 would result in a shadow offset 2 pixels above the text. All burn-in and DVB-Sub font settings must match.
	ShadowYOffset int `json:"shadowYOffset,omitempty" yaml:"shadowYOffset,omitempty"`

	// Specifies the horizontal position of the caption relative to the left side of the output in pixels. A value of 10 would result in the captions starting 10 pixels from the left of the output. If no explicit xPosition is provided, the horizontal caption position will be determined by the alignment parameter. This option is not valid for source captions that are STL, 608/embedded or teletext. These source settings are already pre-defined by the caption stream. All burn-in and DVB-Sub font settings must match.
	XPosition int `json:"xPosition,omitempty" yaml:"xPosition,omitempty"`

	// If no explicit xPosition or yPosition is provided, setting alignment to centered will place the captions at the bottom center of the output. Similarly, setting a left alignment will align captions to the bottom left of the output. If x and y positions are given in conjunction with the alignment parameter, the font will be justified (either left or centered) relative to those coordinates. Selecting “smart” justification will left-justify live subtitles and center-justify pre-recorded subtitles. This option is not valid for source captions that are STL or 608/embedded. These source settings are already pre-defined by the caption stream. All burn-in and DVB-Sub font settings must match.
	Alignment string `json:"alignment,omitempty" yaml:"alignment,omitempty"`

	// External font file used for caption burn-in. File extension must be ‘ttf’ or ‘tte’. Although the user can select output fonts for many different types of input captions, embedded, STL and teletext sources use a strict grid system. Using external fonts with these caption sources could cause unexpected display of proportional fonts. All burn-in and DVB-Sub font settings must match. See Font for more details.
	Font Medialive_ChannelEncoderSettingsCaptionDescriptionDestinationSettingsDvbSubDestinationSettingsFont `json:"font,omitempty" yaml:"font,omitempty"`

	// When set to auto fontSize will scale depending on the size of the output. Giving a positive integer will specify the exact font size in points. All burn-in and DVB-Sub font settings must match.
	FontSize string `json:"fontSize,omitempty" yaml:"fontSize,omitempty"`

	// Specifies the opacity of the background rectangle. 255 is opaque; 0 is transparent. Leaving this parameter blank is equivalent to setting it to 0 (transparent). All burn-in and DVB-Sub font settings must match.
	BackgroundOpacity int `json:"backgroundOpacity,omitempty" yaml:"backgroundOpacity,omitempty"`

	// Specifies font outline color. This option is not valid for source captions that are either 608/embedded or teletext. These source settings are already pre-defined by the caption stream. All burn-in and DVB-Sub font settings must match.
	OutlineColor string `json:"outlineColor,omitempty" yaml:"outlineColor,omitempty"`

	// Specifies font outline size in pixels. This option is not valid for source captions that are either 608/embedded or teletext. These source settings are already pre-defined by the caption stream. All burn-in and DVB-Sub font settings must match.
	OutlineSize int `json:"outlineSize,omitempty" yaml:"outlineSize,omitempty"`

	// Specifies the color of the shadow cast by the captions. All burn-in and DVB-Sub font settings must match.
	ShadowColor string `json:"shadowColor,omitempty" yaml:"shadowColor,omitempty"`

	// Controls whether a fixed grid size will be used to generate the output subtitles bitmap. Only applicable for Teletext inputs and DVB-Sub/Burn-in outputs.
	TeletextGridControl string `json:"teletextGridControl,omitempty" yaml:"teletextGridControl,omitempty"`

	// Specifies the color of the burned-in captions. This option is not valid for source captions that are STL, 608/embedded or teletext. These source settings are already pre-defined by the caption stream. All burn-in and DVB-Sub font settings must match.
	FontColor string `json:"fontColor,omitempty" yaml:"fontColor,omitempty"`

	// Specifies the vertical position of the caption relative to the top of the output in pixels. A value of 10 would result in the captions starting 10 pixels from the top of the output. If no explicit yPosition is provided, the caption will be positioned towards the bottom of the output. This option is not valid for source captions that are STL, 608/embedded or teletext. These source settings are already pre-defined by the caption stream. All burn-in and DVB-Sub font settings must match.
	YPosition int `json:"yPosition,omitempty" yaml:"yPosition,omitempty"`

	// Specifies the opacity of the shadow. 255 is opaque; 0 is transparent. Leaving this parameter blank is equivalent to setting it to 0 (transparent). All burn-in and DVB-Sub font settings must match.
	ShadowOpacity int `json:"shadowOpacity,omitempty" yaml:"shadowOpacity,omitempty"`
}
