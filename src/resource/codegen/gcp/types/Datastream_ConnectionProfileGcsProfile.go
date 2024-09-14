package types

type Datastream_ConnectionProfileGcsProfile struct {
	// The Cloud Storage bucket name.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// The root path inside the Cloud Storage bucket.
	RootPath string `json:"rootPath,omitempty" yaml:"rootPath,omitempty"`
}
