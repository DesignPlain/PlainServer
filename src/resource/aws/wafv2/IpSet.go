package wafv2

type IpSet struct {
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the Region US East (N. Virginia).
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`

	// An array of key:value pairs to associate with the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Contains an array of strings that specifies zero or more IP addresses or blocks of IP addresses. All addresses must be specified using Classless Inter-Domain Routing (CIDR) notation. WAF supports all IPv4 and IPv6 CIDR ranges except for `/0`.
	Addresses []string `json:"addresses,omitempty" yaml:"addresses,omitempty"`

	// A friendly description of the IP set.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Specify IPV4 or IPV6. Valid values are `IPV4` or `IPV6`.
	IpAddressVersion string `json:"ipAddressVersion,omitempty" yaml:"ipAddressVersion,omitempty"`

	// A friendly name of the IP set.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
