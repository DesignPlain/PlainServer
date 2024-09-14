package types

type Comprehend_EntityRecognizerInputDataConfigAnnotations struct {
	// Location of training annotations.
	S3Uri string `json:"s3Uri,omitempty" yaml:"s3Uri,omitempty"`

	//
	TestS3Uri string `json:"testS3Uri,omitempty" yaml:"testS3Uri,omitempty"`
}
