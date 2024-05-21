package bigquery

import types "DesignSphere_Server/src/resource/gcp/types"

type BiReservation struct {
	/*
	   LOCATION_DESCRIPTION


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Preferred tables to use BI capacity for.
	   Structure is documented below.
	*/
	PreferredTables []types.Bigquery_BiReservationPreferredTable `json:"preferredTables,omitempty" yaml:"preferredTables,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// Size of a reservation, in bytes.
	Size int `json:"size,omitempty" yaml:"size,omitempty"`
}
