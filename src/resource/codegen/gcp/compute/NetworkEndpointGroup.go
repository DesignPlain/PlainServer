package compute

type NetworkEndpointGroup struct {
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
	   The network to which all network endpoints in the NEG belong.
	   Uses "default" project network if unspecified.


	   - - -
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   Type of network endpoints in this network endpoint group.
	   NON_GCP_PRIVATE_IP_PORT is used for hybrid connectivity network
	   endpoint groups (see https://cloud.google.com/load-balancing/docs/hybrid).
	   Note that NON_GCP_PRIVATE_IP_PORT can only be used with Backend Services
	   that 1) have the following load balancing schemes: EXTERNAL, EXTERNAL_MANAGED,
	   INTERNAL_MANAGED, and INTERNAL_SELF_MANAGED and 2) support the RATE or
	   CONNECTION balancing modes.
	   Possible values include: GCE_VM_IP, GCE_VM_IP_PORT, NON_GCP_PRIVATE_IP_PORT, INTERNET_IP_PORT, INTERNET_FQDN_PORT, SERVERLESS, and PRIVATE_SERVICE_CONNECT.
	   Default value is `GCE_VM_IP_PORT`.
	   Possible values are: `GCE_VM_IP`, `GCE_VM_IP_PORT`, `NON_GCP_PRIVATE_IP_PORT`, `INTERNET_IP_PORT`, `INTERNET_FQDN_PORT`, `SERVERLESS`, `PRIVATE_SERVICE_CONNECT`.
	*/
	NetworkEndpointType string `json:"networkEndpointType,omitempty" yaml:"networkEndpointType,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Optional subnetwork to which all network endpoints in the NEG belong.
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`

	// Zone where the network endpoint group is located.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
