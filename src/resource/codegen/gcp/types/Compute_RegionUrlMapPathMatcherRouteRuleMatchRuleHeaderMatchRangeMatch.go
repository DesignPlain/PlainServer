package types

type Compute_RegionUrlMapPathMatcherRouteRuleMatchRuleHeaderMatchRangeMatch struct {
	// The end of the range (exclusive).
	RangeEnd int `json:"rangeEnd,omitempty" yaml:"rangeEnd,omitempty"`

	// The start of the range (inclusive).
	RangeStart int `json:"rangeStart,omitempty" yaml:"rangeStart,omitempty"`
}
