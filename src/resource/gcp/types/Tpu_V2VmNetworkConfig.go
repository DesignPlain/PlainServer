package types

type Tpu_V2VmNetworkConfig struct {
	/*
	   The name of the subnetwork for the TPU node. It must be a preexisting Google Compute
	   Engine subnetwork. If both network and subnetwork are specified, the given subnetwork
	   must belong to the given network. If subnetwork is not specified, the subnetwork with the
	   same name as the network will be used.
	*/
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`

	/*
	   Allows the TPU node to send and receive packets with non-matching destination or source
	   IPs. This is required if you plan to use the TPU workers to forward routes.
	*/
	CanIpForward bool `json:"canIpForward,omitempty" yaml:"canIpForward,omitempty"`

	/*
	   Indicates that external IP addresses would be associated with the TPU workers. If set to
	   false, the specified subnetwork or network should have Private Google Access enabled.
	*/
	EnableExternalIps bool `json:"enableExternalIps,omitempty" yaml:"enableExternalIps,omitempty"`

	/*
	   The name of the network for the TPU node. It must be a preexisting Google Compute Engine
	   network. If both network and subnetwork are specified, the given subnetwork must belong
	   to the given network. If network is not specified, it will be looked up from the
	   subnetwork if one is provided, or otherwise use "default".
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`
}
