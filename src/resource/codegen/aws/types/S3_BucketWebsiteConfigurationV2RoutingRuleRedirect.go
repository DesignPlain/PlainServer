package types

type S3_BucketWebsiteConfigurationV2RoutingRuleRedirect struct {
	// Protocol to use when redirecting requests. The default is the protocol that is used in the original request. Valid values: `http`, `https`.
	Protocol string `json:"protocol,omitempty" yaml:"protocol,omitempty"`

	// Object key prefix to use in the redirect request. For example, to redirect requests for all pages with prefix `docs/` (objects in the `docs/` folder) to `documents/`, you can set a `condition` block with `key_prefix_equals` set to `docs/` and in the `redirect` set `replace_key_prefix_with` to `/documents`.
	ReplaceKeyPrefixWith string `json:"replaceKeyPrefixWith,omitempty" yaml:"replaceKeyPrefixWith,omitempty"`

	// Specific object key to use in the redirect request. For example, redirect request to `error.html`.
	ReplaceKeyWith string `json:"replaceKeyWith,omitempty" yaml:"replaceKeyWith,omitempty"`

	// Host name to use in the redirect request.
	HostName string `json:"hostName,omitempty" yaml:"hostName,omitempty"`

	// HTTP redirect code to use on the response.
	HttpRedirectCode string `json:"httpRedirectCode,omitempty" yaml:"httpRedirectCode,omitempty"`
}
