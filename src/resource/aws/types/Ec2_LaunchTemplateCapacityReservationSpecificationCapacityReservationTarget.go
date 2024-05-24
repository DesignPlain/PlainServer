package types

type Ec2_LaunchTemplateCapacityReservationSpecificationCapacityReservationTarget struct {
	// The ID of the Capacity Reservation in which to run the instance.
	CapacityReservationId string `json:"capacityReservationId,omitempty" yaml:"capacityReservationId,omitempty"`

	// The ARN of the Capacity Reservation resource group in which to run the instance.
	CapacityReservationResourceGroupArn string `json:"capacityReservationResourceGroupArn,omitempty" yaml:"capacityReservationResourceGroupArn,omitempty"`
}
