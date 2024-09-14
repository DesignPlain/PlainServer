package types

type Ec2_getNetworkInterfaceAssociation struct {
	// Customer-owned IP address.
	CustomerOwnedIp string `json:"customerOwnedIp,omitempty" yaml:"customerOwnedIp,omitempty"`

	// ID of the Elastic IP address owner.
	IpOwnerId string `json:"ipOwnerId,omitempty" yaml:"ipOwnerId,omitempty"`

	// Public DNS name.
	PublicDnsName string `json:"publicDnsName,omitempty" yaml:"publicDnsName,omitempty"`

	// Address of the Elastic IP address bound to the network interface.
	PublicIp string `json:"publicIp,omitempty" yaml:"publicIp,omitempty"`

	// Allocation ID.
	AllocationId string `json:"allocationId,omitempty" yaml:"allocationId,omitempty"`

	// Association ID.
	AssociationId string `json:"associationId,omitempty" yaml:"associationId,omitempty"`

	// Carrier IP address associated with the network interface. This attribute is only set when the network interface is in a subnet which is associated with a Wavelength Zone.
	CarrierIp string `json:"carrierIp,omitempty" yaml:"carrierIp,omitempty"`
}
