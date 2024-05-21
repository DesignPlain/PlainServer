package types

type Edgecontainer_ClusterMaintenancePolicyWindowRecurringWindow struct {
	/*
	   An RRULE (https://tools.ietf.org/html/rfc5545#section-3.8.5.3) for how
	   this window recurs. They go on for the span of time between the start and
	   end time.
	*/
	Recurrence string `json:"recurrence,omitempty" yaml:"recurrence,omitempty"`

	/*
	   Represents an arbitrary window of time.
	   Structure is documented below.
	*/
	Window Edgecontainer_ClusterMaintenancePolicyWindowRecurringWindowWindow `json:"window,omitempty" yaml:"window,omitempty"`
}
