package types

type Pipes_PipeSourceParametersFilterCriteria struct {
	// An array of up to 5 event patterns. Detailed below.
	Filters []Pipes_PipeSourceParametersFilterCriteriaFilter `json:"filters,omitempty" yaml:"filters,omitempty"`
}
