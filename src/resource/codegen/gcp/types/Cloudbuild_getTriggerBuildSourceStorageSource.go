package types

type Cloudbuild_getTriggerBuildSourceStorageSource struct {
	// Google Cloud Storage bucket containing the source.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	/*
	   Google Cloud Storage generation for the object.
	   If the generation is omitted, the latest generation will be used
	*/
	Generation string `json:"generation,omitempty" yaml:"generation,omitempty"`

	/*
	   Google Cloud Storage object containing the source.
	   This object must be a gzipped archive file (.tar.gz) containing source to build.
	*/
	Object string `json:"object,omitempty" yaml:"object,omitempty"`
}
