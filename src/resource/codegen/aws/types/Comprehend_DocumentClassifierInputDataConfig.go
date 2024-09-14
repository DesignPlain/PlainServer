package types

type Comprehend_DocumentClassifierInputDataConfig struct {
	/*
	   List of training datasets produced by Amazon SageMaker Ground Truth.
	   Used if `data_format` is `AUGMENTED_MANIFEST`.
	   See the `augmented_manifests` Configuration Block section below.
	*/
	AugmentedManifests []Comprehend_DocumentClassifierInputDataConfigAugmentedManifest `json:"augmentedManifests,omitempty" yaml:"augmentedManifests,omitempty"`

	/*
	   The format for the training data.
	   One of `COMPREHEND_CSV` or `AUGMENTED_MANIFEST`.
	*/
	DataFormat string `json:"dataFormat,omitempty" yaml:"dataFormat,omitempty"`

	/*
	   Delimiter between labels when training a multi-label classifier.
	   Valid values are `|`, `~`, `!`, `@`, `#`, `$`, `%!`(MISSING), `^`, `-`, `-`, `_`, `+`, `=`, `\`, `:`, `;`, `>`, `?`, `/`, `<space>`, and `<tab>`.
	   Default is `|`.
	*/
	LabelDelimiter string `json:"labelDelimiter,omitempty" yaml:"labelDelimiter,omitempty"`

	/*
	   Location of training documents.
	   Used if `data_format` is `COMPREHEND_CSV`.
	*/
	S3Uri string `json:"s3Uri,omitempty" yaml:"s3Uri,omitempty"`

	//
	TestS3Uri string `json:"testS3Uri,omitempty" yaml:"testS3Uri,omitempty"`
}
