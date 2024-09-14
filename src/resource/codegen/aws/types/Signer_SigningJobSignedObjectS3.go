package types

type Signer_SigningJobSignedObjectS3 struct {
	// Key name of the object that contains your unsigned code.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	//
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`
}
