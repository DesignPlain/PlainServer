package ec2

type CapacityReservation struct {
	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
	Tenancy string `json:"tenancy,omitempty" yaml:"tenancy,omitempty"`

	// The Availability Zone in which to create the Capacity Reservation.
	AvailabilityZone string `json:"availabilityZone,omitempty" yaml:"availabilityZone,omitempty"`

	// Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
	EphemeralStorage bool `json:"ephemeralStorage,omitempty" yaml:"ephemeralStorage,omitempty"`

	// The number of instances for which to reserve capacity.
	InstanceCount int `json:"instanceCount,omitempty" yaml:"instanceCount,omitempty"`

	// Indicates the type of instance launches that the Capacity Reservation accepts. Specify either `open` or `targeted`.
	InstanceMatchCriteria string `json:"instanceMatchCriteria,omitempty" yaml:"instanceMatchCriteria,omitempty"`

	// The Amazon Resource Name (ARN) of the Outpost on which to create the Capacity Reservation.
	OutpostArn string `json:"outpostArn,omitempty" yaml:"outpostArn,omitempty"`

	// The Amazon Resource Name (ARN) of the cluster placement group in which to create the Capacity Reservation.
	PlacementGroupArn string `json:"placementGroupArn,omitempty" yaml:"placementGroupArn,omitempty"`

	// Indicates whether the Capacity Reservation supports EBS-optimized instances.
	EbsOptimized bool `json:"ebsOptimized,omitempty" yaml:"ebsOptimized,omitempty"`

	// The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
	EndDate string `json:"endDate,omitempty" yaml:"endDate,omitempty"`

	// Indicates the way in which the Capacity Reservation ends. Specify either `unlimited` or `limited`.
	EndDateType string `json:"endDateType,omitempty" yaml:"endDateType,omitempty"`

	// The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
	InstancePlatform string `json:"instancePlatform,omitempty" yaml:"instancePlatform,omitempty"`

	// The instance type for which to reserve capacity.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`
}
