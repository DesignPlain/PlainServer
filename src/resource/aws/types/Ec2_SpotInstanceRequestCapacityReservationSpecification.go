package types

type Ec2_SpotInstanceRequestCapacityReservationSpecification struct {
	// Indicates the instance's Capacity Reservation preferences. Can be `"open"` or `"none"`. (Default: `"open"`).
	CapacityReservationPreference string `json:"capacityReservationPreference,omitempty" yaml:"capacityReservationPreference,omitempty"`

	/*
	   Information about the target Capacity Reservation. See Capacity Reservation Target below for more details.

	   For more information, see the documentation on [Capacity Reservations](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/capacity-reservations-using.html).
	*/
	CapacityReservationTarget Ec2_SpotInstanceRequestCapacityReservationSpecificationCapacityReservationTarget `json:"capacityReservationTarget,omitempty" yaml:"capacityReservationTarget,omitempty"`
}
