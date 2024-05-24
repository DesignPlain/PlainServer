package types

type Ec2_SpotInstanceRequestNetworkInterface struct {
	// Integer index of the network interface attachment. Limited by instance type.
	DeviceIndex int `json:"deviceIndex,omitempty" yaml:"deviceIndex,omitempty"`

	// Integer index of the network card. Limited by instance type. The default index is `0`.
	NetworkCardIndex int `json:"networkCardIndex,omitempty" yaml:"networkCardIndex,omitempty"`

	// ID of the network interface to attach.
	NetworkInterfaceId string `json:"networkInterfaceId,omitempty" yaml:"networkInterfaceId,omitempty"`

	// Whether or not to delete the network interface on instance termination. Defaults to `false`. Currently, the only valid value is `false`, as this is only supported when creating new network interfaces when launching an instance.
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" yaml:"deleteOnTermination,omitempty"`
}
