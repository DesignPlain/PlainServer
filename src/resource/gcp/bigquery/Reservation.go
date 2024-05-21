package bigquery

import types "DesignSphere_Server/src/resource/gcp/types"

type Reservation struct {
	/*
	   If false, any query using this reservation will use idle slots from other reservations within
	   the same admin project. If true, a query using this reservation will execute with the slot
	   capacity specified above at most.
	*/
	IgnoreIdleSlots bool `json:"ignoreIdleSlots,omitempty" yaml:"ignoreIdleSlots,omitempty"`

	/*
	   Applicable only for reservations located within one of the BigQuery multi-regions (US or EU).
	   If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region.
	*/
	MultiRegionAuxiliary bool `json:"multiRegionAuxiliary,omitempty" yaml:"multiRegionAuxiliary,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The configuration parameters for the auto scaling feature.
	   Structure is documented below.
	*/
	Autoscale types.Bigquery_ReservationAutoscale `json:"autoscale,omitempty" yaml:"autoscale,omitempty"`

	// The edition type. Valid values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS
	Edition string `json:"edition,omitempty" yaml:"edition,omitempty"`

	/*
	   The name of the reservation. This field must only contain alphanumeric characters or dash.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
	   unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	*/
	SlotCapacity int `json:"slotCapacity,omitempty" yaml:"slotCapacity,omitempty"`

	// Maximum number of queries that are allowed to run concurrently in this reservation. This is a soft limit due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency will be automatically set based on the reservation size.
	Concurrency int `json:"concurrency,omitempty" yaml:"concurrency,omitempty"`

	/*
	   The geographic location where the transfer config should reside.
	   Examples: US, EU, asia-northeast1. The default value is US.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`
}
