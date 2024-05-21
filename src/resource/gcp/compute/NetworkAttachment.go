package compute

type NetworkAttachment struct {
	// Projects that are not allowed to connect to this network attachment. The project can be specified using its id or number.
	ProducerRejectLists []string `json:"producerRejectLists,omitempty" yaml:"producerRejectLists,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   URL of the region where the network attachment resides. This field applies only to the region resource. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.


	   - - -
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// An array of URLs where each entry is the URL of a subnet provided by the service consumer to use for endpoints in the producers that connect to this network attachment.
	Subnetworks []string `json:"subnetworks,omitempty" yaml:"subnetworks,omitempty"`

	/*
	   The connection preference of service attachment. The value can be set to ACCEPT_AUTOMATIC. An ACCEPT_AUTOMATIC service attachment is one that always accepts the connection from consumer forwarding rules.
	   Possible values are: `ACCEPT_AUTOMATIC`, `ACCEPT_MANUAL`, `INVALID`.
	*/
	ConnectionPreference string `json:"connectionPreference,omitempty" yaml:"connectionPreference,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Projects that are allowed to connect to this network attachment. The project can be specified using its id or number.
	ProducerAcceptLists []string `json:"producerAcceptLists,omitempty" yaml:"producerAcceptLists,omitempty"`
}
