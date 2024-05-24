package types

type Comprehend_EntityRecognizerInputDataConfig struct {
	/*
	   Specifies location of the document annotation data.
	   See the `annotations` Configuration Block section below.
	   One of `annotations` or `entity_list` is required.
	*/
	Annotations Comprehend_EntityRecognizerInputDataConfigAnnotations `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	/*
	   List of training datasets produced by Amazon SageMaker Ground Truth.
	   Used if `data_format` is `AUGMENTED_MANIFEST`.
	   See the `augmented_manifests` Configuration Block section below.
	*/
	AugmentedManifests []Comprehend_EntityRecognizerInputDataConfigAugmentedManifest `json:"augmentedManifests,omitempty" yaml:"augmentedManifests,omitempty"`

	/*
	   The format for the training data.
	   One of `COMPREHEND_CSV` or `AUGMENTED_MANIFEST`.
	*/
	DataFormat string `json:"dataFormat,omitempty" yaml:"dataFormat,omitempty"`

	/*
	   Specifies a collection of training documents.
	   Used if `data_format` is `COMPREHEND_CSV`.
	   See the `documents` Configuration Block section below.
	*/
	Documents Comprehend_EntityRecognizerInputDataConfigDocuments `json:"documents,omitempty" yaml:"documents,omitempty"`

	/*
	   Specifies location of the entity list data.
	   See the `entity_list` Configuration Block section below.
	   One of `entity_list` or `annotations` is required.
	*/
	EntityList Comprehend_EntityRecognizerInputDataConfigEntityList `json:"entityList,omitempty" yaml:"entityList,omitempty"`

	/*
	   Set of entity types to be recognized.
	   Has a maximum of 25 items.
	   See the `entity_types` Configuration Block section below.
	*/
	EntityTypes []Comprehend_EntityRecognizerInputDataConfigEntityType `json:"entityTypes,omitempty" yaml:"entityTypes,omitempty"`
}
