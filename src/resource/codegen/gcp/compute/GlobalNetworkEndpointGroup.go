package compute

type GlobalNetworkEndpointGroup struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The default port used if the port number is not specified in the
	   network endpoint.
	*/
	DefaultPort int `json:"defaultPort,omitempty" yaml:"defaultPort,omitempty"`

	/*
	   An optional description of this resource. Provide this property when
	   you create the resource.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Name of the resource; provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035. Specifically, the name must be 1-63 characters long and match
	   the regular expression `a-z?` which means the
	   first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Type of network endpoints in this network endpoint group.
	   Possible values are: `INTERNET_IP_PORT`, `INTERNET_FQDN_PORT`.


	   - - -
	*/
	NetworkEndpointType string `json:"networkEndpointType,omitempty" yaml:"networkEndpointType,omitempty"`
}
