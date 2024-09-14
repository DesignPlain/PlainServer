package types

type S3_BucketWebsiteConfigurationV2IndexDocument struct {
	/*
	   Suffix that is appended to a request that is for a directory on the website endpoint.
	   For example, if the suffix is `index.html` and you make a request to `samplebucket/images/`, the data that is returned will be for the object with the key name `images/index.html`.
	   The suffix must not be empty and must not include a slash character.
	*/
	Suffix string `json:"suffix,omitempty" yaml:"suffix,omitempty"`
}
