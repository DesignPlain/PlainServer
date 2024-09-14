package types

type Comprehend_EntityRecognizerInputDataConfigDocuments struct {
	// Location of training documents.
	S3Uri string `json:"s3Uri,omitempty" yaml:"s3Uri,omitempty"`

	//
	TestS3Uri string `json:"testS3Uri,omitempty" yaml:"testS3Uri,omitempty"`

	/*
	   Specifies how the input files should be processed.
	   One of `ONE_DOC_PER_LINE` or `ONE_DOC_PER_FILE`.
	*/
	InputFormat string `json:"inputFormat,omitempty" yaml:"inputFormat,omitempty"`
}
