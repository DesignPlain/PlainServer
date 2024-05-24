package types

type Signer_SigningJobSourceS3 struct {
	// Name of the S3 bucket.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Key name of the object that contains your unsigned code.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Version of your source image in your version enabled S3 bucket.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
