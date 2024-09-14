package wafregional

import types "libds/aws/types"

type ByteMatchSet struct {
	// Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
	ByteMatchTuples []types.Wafregional_ByteMatchSetByteMatchTuple `json:"byteMatchTuples,omitempty" yaml:"byteMatchTuples,omitempty"`

	// The name or description of the ByteMatchSet.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
