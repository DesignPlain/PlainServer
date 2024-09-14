package types

type Ec2_SpotInstanceRequestCapacityReservationSpecificationCapacityReservationTarget struct {
	// ID of the Capacity Reservation in which to run the instance.
	CapacityReservationId string `json:"capacityReservationId,omitempty" yaml:"capacityReservationId,omitempty"`

	// ARN of the Capacity Reservation resource group in which to run the instance.
	CapacityReservationResourceGroupArn string `json:"capacityReservationResourceGroupArn,omitempty" yaml:"capacityReservationResourceGroupArn,omitempty"`
}
