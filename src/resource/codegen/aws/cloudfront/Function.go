package cloudfront

type Function struct {
	// Source code of the function
	Code string `json:"code,omitempty" yaml:"code,omitempty"`

	// Comment.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	// List of `aws.cloudfront.KeyValueStore` ARNs to be associated to the function. AWS limits associations to on key value store per function.
	KeyValueStoreAssociations []string `json:"keyValueStoreAssociations,omitempty" yaml:"keyValueStoreAssociations,omitempty"`

	// Unique name for your CloudFront Function.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Whether to publish creation/change as Live CloudFront Function Version. Defaults to `true`.
	Publish bool `json:"publish,omitempty" yaml:"publish,omitempty"`

	/*
	   Identifier of the function's runtime. Valid values are `cloudfront-js-1.0` and `cloudfront-js-2.0`.

	   The following arguments are optional:
	*/
	Runtime string `json:"runtime,omitempty" yaml:"runtime,omitempty"`
}
