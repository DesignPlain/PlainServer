package types

type Lex_V2modelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureResponseMessageGroupVariationImageResponseCard struct {
	// Configuration blocks for buttons that should be displayed on the response card. The arrangement of the buttons is determined by the platform that displays the button. See `button`.
	Buttons []Lex_V2modelsIntentInitialResponseSettingCodeHookPostCodeHookSpecificationFailureResponseMessageGroupVariationImageResponseCardButton `json:"buttons,omitempty" yaml:"buttons,omitempty"`

	// URL of an image to display on the response card. The image URL must be publicly available so that the platform displaying the response card has access to the image.
	ImageUrl string `json:"imageUrl,omitempty" yaml:"imageUrl,omitempty"`

	// Subtitle to display on the response card. The format of the subtitle is determined by the platform displaying the response card.
	Subtitle string `json:"subtitle,omitempty" yaml:"subtitle,omitempty"`

	// Title to display on the response card. The format of the title is determined by the platform displaying the response card.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
}
