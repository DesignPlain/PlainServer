package compute

type NetworkEndpoint struct {
	/*
	   The name for a specific VM instance that the IP address belongs to.
	   This is required for network endpoints of type GCE_VM_IP_PORT.
	   The instance must be in the same zone of network endpoint group.
	*/
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`

	/*
	   IPv4 address of network endpoint. The IP address must belong
	   to a VM in GCE (either the primary IP or as part of an aliased IP
	   range).
	*/
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`

	/*
	   The network endpoint group this endpoint is part of.


	   - - -
	*/
	NetworkEndpointGroup string `json:"networkEndpointGroup,omitempty" yaml:"networkEndpointGroup,omitempty"`

	/*
	   Port number of network endpoint.
	   --Note-- `port` is required unless the Network Endpoint Group is created
	   with the type of `GCE_VM_IP`
	*/
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Zone where the containing network endpoint group is located.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
