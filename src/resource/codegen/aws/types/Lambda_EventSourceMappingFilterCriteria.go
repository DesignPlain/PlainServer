package types

type Lambda_EventSourceMappingFilterCriteria struct {
	// A set of up to 5 filter. If an event satisfies at least one, Lambda sends the event to the function or adds it to the next batch. Detailed below.
	Filters []Lambda_EventSourceMappingFilterCriteriaFilter `json:"filters,omitempty" yaml:"filters,omitempty"`
}
