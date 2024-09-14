package compute

import types "libds/gcp/types"

type Reservation struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The share setting for reservations.
	   Structure is documented below.
	*/
	ShareSettings types.Compute_ReservationShareSettings `json:"shareSettings,omitempty" yaml:"shareSettings,omitempty"`

	/*
	   Reservation for instances with specific machine shapes.
	   Structure is documented below.
	*/
	SpecificReservation types.Compute_ReservationSpecificReservation `json:"specificReservation,omitempty" yaml:"specificReservation,omitempty"`

	/*
	   When set to true, only VMs that target this reservation by name can
	   consume this reservation. Otherwise, it can be consumed by VMs with
	   affinity for any reservation. Defaults to false.
	*/
	SpecificReservationRequired bool `json:"specificReservationRequired,omitempty" yaml:"specificReservationRequired,omitempty"`

	// The zone where the reservation is made.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`

	// An optional description of this resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Name of the resource. Provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035. Specifically, the name must be 1-63 characters long and match
	   the regular expression `a-z?` which means the
	   first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the last
	   character, which cannot be a dash.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
