package compute

type GlobalAddress struct {
	/*
	   The type of the address to reserve.
	   - EXTERNAL indicates public/external single IP address.
	   - INTERNAL indicates internal IP ranges belonging to some network.
	   Default value is `EXTERNAL`.
	   Possible values are: `EXTERNAL`, `INTERNAL`.
	*/
	AddressType string `json:"addressType,omitempty" yaml:"addressType,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The IP Version that will be used by this address. The default value is `IPV4`.
	   Possible values are: `IPV4`, `IPV6`.
	*/
	IpVersion string `json:"ipVersion,omitempty" yaml:"ipVersion,omitempty"`

	/*
	   Labels to apply to this address.  A list of key->value pairs.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The purpose of the resource. Possible values include:
	   - VPC_PEERING - for peer networks
	   - PRIVATE_SERVICE_CONNECT - for  Private Service Connect networks
	*/
	Purpose string `json:"purpose,omitempty" yaml:"purpose,omitempty"`

	/*
	   The IP address or beginning of the address range represented by this
	   resource. This can be supplied as an input to reserve a specific
	   address or omitted to allow GCP to choose a valid one for you.
	*/
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	/*
	   Name of the resource. Provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035.  Specifically, the name must be 1-63 characters long and
	   match the regular expression `a-z?` which means
	   the first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The URL of the network in which to reserve the IP range. The IP range
	   must be in RFC1918 space. The network cannot be deleted if there are
	   any reserved IP ranges referring to it.
	   This should only be set when using an Internal address.
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The prefix length of the IP range. If not present, it means the
	   address field is a single IP address.
	   This field is not applicable to addresses with addressType=INTERNAL
	   when purpose=PRIVATE_SERVICE_CONNECT
	*/
	PrefixLength int `json:"prefixLength,omitempty" yaml:"prefixLength,omitempty"`
}
