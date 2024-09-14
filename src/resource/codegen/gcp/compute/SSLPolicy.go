package compute

type SSLPolicy struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Profile specifies the set of SSL features that can be used by the
	   load balancer when negotiating SSL with clients. This can be one of
	   `COMPATIBLE`, `MODERN`, `RESTRICTED`, or `CUSTOM`. If using `CUSTOM`,
	   the set of SSL features to enable must be specified in the
	   `customFeatures` field.
	   See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	   for which ciphers are available to use. --Note--: this argument
	   -must- be present when using the `CUSTOM` profile. This argument
	   -must not- be present when using any other profile.
	*/
	CustomFeatures []string `json:"customFeatures,omitempty" yaml:"customFeatures,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The minimum version of SSL protocol that can be used by the clients
	   to establish a connection with the load balancer.
	   Default value is `TLS_1_0`.
	   Possible values are: `TLS_1_0`, `TLS_1_1`, `TLS_1_2`.
	*/
	MinTlsVersion string `json:"minTlsVersion,omitempty" yaml:"minTlsVersion,omitempty"`

	/*
	   Name of the resource. Provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035. Specifically, the name must be 1-63 characters long and match
	   the regular expression `a-z?` which means the
	   first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Profile specifies the set of SSL features that can be used by the
	   load balancer when negotiating SSL with clients. If using `CUSTOM`,
	   the set of SSL features to enable must be specified in the
	   `customFeatures` field.
	   See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
	   for information on what cipher suites each profile provides. If
	   `CUSTOM` is used, the `custom_features` attribute --must be set--.
	   Default value is `COMPATIBLE`.
	   Possible values are: `COMPATIBLE`, `MODERN`, `RESTRICTED`, `CUSTOM`.
	*/
	Profile string `json:"profile,omitempty" yaml:"profile,omitempty"`
}
