package types

type Ec2_LaunchTemplateCapacityReservationSpecification struct {
	// Indicates the instance's Capacity Reservation preferences. Can be `open` or `none`. (Default `none`).
	CapacityReservationPreference string `json:"capacityReservationPreference,omitempty" yaml:"capacityReservationPreference,omitempty"`

	// Used to target a specific Capacity Reservation:
	CapacityReservationTarget Ec2_LaunchTemplateCapacityReservationSpecificationCapacityReservationTarget `json:"capacityReservationTarget,omitempty" yaml:"capacityReservationTarget,omitempty"`
}
