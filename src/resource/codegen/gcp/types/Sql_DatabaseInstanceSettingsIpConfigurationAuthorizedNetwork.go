package types

type Sql_DatabaseInstanceSettingsIpConfigurationAuthorizedNetwork struct {
	/*
	   The [RFC 3339](https://tools.ietf.org/html/rfc3339)
	   formatted date time string indicating when this whitelist expires.
	*/
	ExpirationTime string `json:"expirationTime,omitempty" yaml:"expirationTime,omitempty"`

	// A name for this whitelist entry.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   A CIDR notation IPv4 or IPv6 address that is allowed to
	   access this instance. Must be set even if other two attributes are not for
	   the whitelist to become active.
	*/
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
