package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationCharacterMaskConfigCharactersToIgnore struct {
	/*
	   Common characters to not transform when masking. Useful to avoid removing punctuation. Only one of this or `characters_to_skip` must be specified.
	   Possible values are: `NUMERIC`, `ALPHA_UPPER_CASE`, `ALPHA_LOWER_CASE`, `PUNCTUATION`, `WHITESPACE`.
	*/
	CommonCharactersToIgnore string `json:"commonCharactersToIgnore,omitempty" yaml:"commonCharactersToIgnore,omitempty"`

	// Characters to not transform when masking. Only one of this or `common_characters_to_ignore` must be specified.
	CharactersToSkip string `json:"charactersToSkip,omitempty" yaml:"charactersToSkip,omitempty"`
}
