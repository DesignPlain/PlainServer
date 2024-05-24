package wafv2

import types "DesignSphere_Server/src/resource/aws/types"

type RegexPatternSet struct {
	// A friendly description of the regular expression pattern set.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A friendly name of the regular expression pattern set.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details. A maximum of 10 `regular_expression` blocks may be specified.
	RegularExpressions []types.Wafv2_RegexPatternSetRegularExpression `json:"regularExpressions,omitempty" yaml:"regularExpressions,omitempty"`

	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`

	// An array of key:value pairs to associate with the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
