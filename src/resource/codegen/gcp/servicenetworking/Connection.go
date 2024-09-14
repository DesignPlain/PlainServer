package servicenetworking

type Connection struct {
	/*
	   When set to ABANDON, terraform will abandon management of the resource instead of deleting it. Prevents terraform apply
	   failures with CloudSQL. Note: The resource will still exist.
	*/
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`

	// Name of VPC network connected with service producers using VPC peering.
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   Named IP address range(s) of PEERING type reserved for
	   this service provider. Note that invoking this method with a different range when connection
	   is already established will not reallocate already provisioned service producer subnetworks.
	*/
	ReservedPeeringRanges []string `json:"reservedPeeringRanges,omitempty" yaml:"reservedPeeringRanges,omitempty"`

	/*
	   Provider peering service that is managing peering connectivity for a
	   service provider organization. For Google services that support this functionality it is
	   'servicenetworking.googleapis.com'.
	*/
	Service string `json:"service,omitempty" yaml:"service,omitempty"`
}
