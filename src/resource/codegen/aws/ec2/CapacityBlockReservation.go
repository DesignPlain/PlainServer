package ec2

import types "libds/aws/types"

type CapacityBlockReservation struct {
	// The Capacity Block Reservation ID.
	CapacityBlockOfferingId string `json:"capacityBlockOfferingId,omitempty" yaml:"capacityBlockOfferingId,omitempty"`

	// The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
	InstancePlatform string `json:"instancePlatform,omitempty" yaml:"instancePlatform,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	//
	Timeouts types.Ec2_CapacityBlockReservationTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`
}
