package cloudfront

import types "libds/aws/types"

type KeyValueStore struct {
	// Comment.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	/*
	   Unique name for your CloudFront KeyValueStore.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Timeouts types.Cloudfront_KeyValueStoreTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
