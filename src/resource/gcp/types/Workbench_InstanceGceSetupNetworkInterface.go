package types

type Workbench_InstanceGceSetupNetworkInterface struct {
	// Optional. The name of the VPC that this VM instance is in.
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   Optional. The type of vNIC to be used on this interface. This
	   may be gVNIC or VirtioNet.
	   Possible values are: `VIRTIO_NET`, `GVNIC`.
	*/
	NicType string `json:"nicType,omitempty" yaml:"nicType,omitempty"`

	// Optional. The name of the subnet that this VM instance is in.
	Subnet string `json:"subnet,omitempty" yaml:"subnet,omitempty"`
}
