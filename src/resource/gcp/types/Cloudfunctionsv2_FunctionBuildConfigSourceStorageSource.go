package types

type Cloudfunctionsv2_FunctionBuildConfigSourceStorageSource struct {
	/*
	   Google Cloud Storage generation for the object. If the generation
	   is omitted, the latest generation will be used.
	*/
	Generation int `json:"generation,omitempty" yaml:"generation,omitempty"`

	// Google Cloud Storage object containing the source.
	Object string `json:"object,omitempty" yaml:"object,omitempty"`

	// Google Cloud Storage bucket containing the source
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`
}
