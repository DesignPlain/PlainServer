package types

type Ec2_getLaunchTemplateCapacityReservationSpecification struct {
	//
	CapacityReservationPreference string `json:"capacityReservationPreference,omitempty" yaml:"capacityReservationPreference,omitempty"`

	//
	CapacityReservationTargets []Ec2_getLaunchTemplateCapacityReservationSpecificationCapacityReservationTarget `json:"capacityReservationTargets,omitempty" yaml:"capacityReservationTargets,omitempty"`
}
