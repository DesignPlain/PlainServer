package types

type S3_BucketWebsiteConfigurationV2RoutingRuleCondition struct {
	// HTTP error code when the redirect is applied. If specified with `key_prefix_equals`, then both must be true for the redirect to be applied.
	HttpErrorCodeReturnedEquals string `json:"httpErrorCodeReturnedEquals,omitempty" yaml:"httpErrorCodeReturnedEquals,omitempty"`

	// Object key name prefix when the redirect is applied. If specified with `http_error_code_returned_equals`, then both must be true for the redirect to be applied.
	KeyPrefixEquals string `json:"keyPrefixEquals,omitempty" yaml:"keyPrefixEquals,omitempty"`
}
