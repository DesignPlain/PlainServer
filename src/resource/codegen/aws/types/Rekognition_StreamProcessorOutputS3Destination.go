package types

type Rekognition_StreamProcessorOutputS3Destination struct {
	// Name of the Amazon S3 bucket you want to associate with the streaming video project.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// The prefix value of the location within the bucket that you want the information to be published to.
	KeyPrefix string `json:"keyPrefix,omitempty" yaml:"keyPrefix,omitempty"`
}
