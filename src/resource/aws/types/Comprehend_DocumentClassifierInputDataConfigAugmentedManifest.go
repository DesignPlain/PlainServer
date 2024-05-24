package types

type Comprehend_DocumentClassifierInputDataConfigAugmentedManifest struct {
	// Location of annotation files.
	AnnotationDataS3Uri string `json:"annotationDataS3Uri,omitempty" yaml:"annotationDataS3Uri,omitempty"`

	// The JSON attribute that contains the annotations for the training documents.
	AttributeNames []string `json:"attributeNames,omitempty" yaml:"attributeNames,omitempty"`

	/*
	   Type of augmented manifest.
	   One of `PLAIN_TEXT_DOCUMENT` or `SEMI_STRUCTURED_DOCUMENT`.
	*/
	DocumentType string `json:"documentType,omitempty" yaml:"documentType,omitempty"`

	// Location of augmented manifest file.
	S3Uri string `json:"s3Uri,omitempty" yaml:"s3Uri,omitempty"`

	// Location of source PDF files.
	SourceDocumentsS3Uri string `json:"sourceDocumentsS3Uri,omitempty" yaml:"sourceDocumentsS3Uri,omitempty"`

	/*
	   Purpose of data in augmented manifest.
	   One of `TRAIN` or `TEST`.
	*/
	Split string `json:"split,omitempty" yaml:"split,omitempty"`
}
