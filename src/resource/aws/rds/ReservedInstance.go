package rds

type ReservedInstance struct {
	// Number of instances to reserve. Default value is `1`.
	InstanceCount int `json:"instanceCount,omitempty" yaml:"instanceCount,omitempty"`

	/*
	   ID of the Reserved DB instance offering to purchase. To determine an `offering_id`, see the `aws.rds.getReservedInstanceOffering` data source.

	   The following arguments are optional:
	*/
	OfferingId string `json:"offeringId,omitempty" yaml:"offeringId,omitempty"`

	// Customer-specified identifier to track this reservation.
	ReservationId string `json:"reservationId,omitempty" yaml:"reservationId,omitempty"`

	// Map of tags to assign to the DB reservation. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
