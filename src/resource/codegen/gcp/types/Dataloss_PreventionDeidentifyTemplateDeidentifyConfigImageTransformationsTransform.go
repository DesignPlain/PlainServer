package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransform struct {
	/*
	   Apply transformation to the selected infoTypes.
	   Structure is documented below.
	*/
	SelectedInfoTypes Dataloss_PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformSelectedInfoTypes `json:"selectedInfoTypes,omitempty" yaml:"selectedInfoTypes,omitempty"`

	// Apply transformation to all findings not specified in other ImageTransformation's selectedInfoTypes.
	AllInfoTypes Dataloss_PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformAllInfoTypes `json:"allInfoTypes,omitempty" yaml:"allInfoTypes,omitempty"`

	// Apply transformation to all text that doesn't match an infoType.
	AllText Dataloss_PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformAllText `json:"allText,omitempty" yaml:"allText,omitempty"`

	/*
	   The color to use when redacting content from an image. If not specified, the default is black.
	   Structure is documented below.
	*/
	RedactionColor Dataloss_PreventionDeidentifyTemplateDeidentifyConfigImageTransformationsTransformRedactionColor `json:"redactionColor,omitempty" yaml:"redactionColor,omitempty"`
}
