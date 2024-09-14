package ec2

type SubnetCidrReservation struct {
	// The ID of the subnet to create the reservation for.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	// The CIDR block for the reservation.
	CidrBlock string `json:"cidrBlock,omitempty" yaml:"cidrBlock,omitempty"`

	// A brief description of the reservation.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The type of reservation to create. Valid values: `explicit`, `prefix`
	ReservationType string `json:"reservationType,omitempty" yaml:"reservationType,omitempty"`
}
