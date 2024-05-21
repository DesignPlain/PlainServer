package types

type Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceDictionaryConfig struct {
	/*
	   A list of words to select from for random replacement. The [limits](https://cloud.google.com/dlp/limits) page contains details about the size limits of dictionaries.
	   Structure is documented below.
	*/
	WordList Dataloss_PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationPrimitiveTransformationReplaceDictionaryConfigWordList `json:"wordList,omitempty" yaml:"wordList,omitempty"`
}
