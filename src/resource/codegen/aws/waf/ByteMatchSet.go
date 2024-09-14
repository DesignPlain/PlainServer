package waf

import types "libds/aws/types"

type ByteMatchSet struct {
	// The name or description of the Byte Match Set.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Specifies the bytes (typically a string that corresponds
	   with ASCII characters) that you want to search for in web requests,
	   the location in requests that you want to search, and other settings.
	*/
	ByteMatchTuples []types.Waf_ByteMatchSetByteMatchTuple `json:"byteMatchTuples,omitempty" yaml:"byteMatchTuples,omitempty"`
}
