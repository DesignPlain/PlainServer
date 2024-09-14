package ec2

type EipAssociation struct {
	// The allocation ID. This is required for EC2-VPC.
	AllocationId string `json:"allocationId,omitempty" yaml:"allocationId,omitempty"`

	/*
	   Whether to allow an Elastic IP to
	   be re-associated. Defaults to `true` in VPC.
	*/
	AllowReassociation bool `json:"allowReassociation,omitempty" yaml:"allowReassociation,omitempty"`

	/*
	   The ID of the instance. This is required for
	   EC2-Classic. For EC2-VPC, you can specify either the instance ID or the
	   network interface ID, but not both. The operation fails if you specify an
	   instance ID unless exactly one network interface is attached.
	*/
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	/*
	   The ID of the network interface. If the
	   instance has more than one network interface, you must specify a network
	   interface ID.
	*/
	NetworkInterfaceId string `json:"networkInterfaceId,omitempty" yaml:"networkInterfaceId,omitempty"`

	/*
	   The primary or secondary private IP address
	   to associate with the Elastic IP address. If no private IP address is
	   specified, the Elastic IP address is associated with the primary private IP
	   address.
	*/
	PrivateIpAddress string `json:"privateIpAddress,omitempty" yaml:"privateIpAddress,omitempty"`

	// The Elastic IP address. This is required for EC2-Classic.
	PublicIp string `json:"publicIp,omitempty" yaml:"publicIp,omitempty"`
}
