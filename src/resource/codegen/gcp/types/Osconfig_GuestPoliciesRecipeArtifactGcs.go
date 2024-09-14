package types

type Osconfig_GuestPoliciesRecipeArtifactGcs struct {
	/*
	   Bucket of the Google Cloud Storage object. Given an example URL: https://storage.googleapis.com/my-bucket/foo/bar#1234567
	   this value would be my-bucket.
	*/
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	/*
	   Must be provided if allowInsecure is false. Generation number of the Google Cloud Storage object.
	   https://storage.googleapis.com/my-bucket/foo/bar#1234567 this value would be 1234567.
	*/
	Generation int `json:"generation,omitempty" yaml:"generation,omitempty"`

	/*
	   Name of the Google Cloud Storage object. Given an example URL: https://storage.googleapis.com/my-bucket/foo/bar#1234567
	   this value would be foo/bar.
	*/
	Object string `json:"object,omitempty" yaml:"object,omitempty"`
}
